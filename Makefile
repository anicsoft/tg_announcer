build:
	go build -o ./build ./cmd/companies_service
	go build -o ./build ./cmd/tg_bot

run-service: build
	./build/companies_service

run-bot: build
	./build/tg_bot

.PHONY: all build run-service run-bot