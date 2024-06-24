LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest

get-deps:
	go get github.com/swaggo/swag
	go get -u github.com/swaggo/http-swagger

swag:
	${LOCAL_BIN}/swag init -g cmd/application/main.go

swag-fmt:
	${LOCAL_BIN}/swag fmt

build:
	go build -o ./build/ ./cmd/application

run: build
	./build/application

.PHONY: all build run