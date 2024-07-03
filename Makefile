include .env
LOCAL_BIN:=$(CURDIR)/backend/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest

local-migration-status:
	${LOCAL_BIN}/goose -dir migrations postgres "host=${POSTGRES_HOST} port=${POSTGRES_PORT} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" status -v

local-migration-up:
	${LOCAL_BIN}/goose -dir migrations postgres "host=${POSTGRES_HOST} port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" up -v

local-migration-down:
	${LOCAL_BIN}/goose -dir migrations postgres "host=${POSTGRES_HOST} port=${POSTGRES_PORT} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" down -v
