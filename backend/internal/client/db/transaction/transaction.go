package transaction

import (
	"context"
	"tg_announcer/internal/client/db"
	"tg_announcer/internal/client/db/pg"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type manager struct {
	db db.Transactor
}

func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{db: db}
}

func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}

func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	ctx = pg.MakeContextTx(ctx, tx)

	defer func() {
		if p := recover(); p != nil {
			err = errors.Errorf("panic recovered: %v", p)
		}

		if err != nil {
			if errRollBack := tx.Rollback(ctx); errRollBack != nil {
				err = errors.Wrap(errRollBack, "failed to roll back transaction")
			}

			return
		}

		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "failed to commit transaction")
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return nil
}
