build:
	go build -o ./build ./cmd/companies_service

run: build
	./build/companies_service

.PHONY: all build run