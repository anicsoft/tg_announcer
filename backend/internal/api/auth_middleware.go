package api

import (
	"context"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	_initDataKey contextKey = "init-data"
)

type contextKey string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We expect passing init data in the Authorization header in the following format:
		// <auth-type> <auth-data>
		// <auth-type> must be "tma", and <auth-data> is Telegram Mini Apps init data.
		authHeader := r.Header.Get("Authorization")
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		authType := authParts[0]
		authData := authParts[1]

		switch authType {
		case "tma":
			// Validate init data. We consider init data sign valid for 1 hour from their creation moment.
			if err := initdata.Validate(authData, os.Getenv("TELEGRAM_BOT_TOKEN"), time.Hour); err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			// Parse init data. We will surely need it in the future.
			initData, err := initdata.Parse(authData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Add init data to the request context
			ctx := withInitData(r.Context(), initData)
			r = r.WithContext(ctx)
		default:
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func withInitData(ctx context.Context, initData any) context.Context {
	// Assuming you have a type for init data, replace `any` with the correct type
	return context.WithValue(ctx, _initDataKey, initData)
}

func ctxInitData(ctx context.Context) (initdata.InitData, bool) {
	initData, ok := ctx.Value(_initDataKey).(initdata.InitData)
	return initData, ok
}
