#!/usr/bin/env bash
source .env

export MIGRATION_DSN="host=$POSTGRES_HOST port=$POSTGRES_PORT dbname=$POSTGRES_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"

sleep 2 && goose -dir "migrations" postgres "${MIGRATION_DSN}" up -v