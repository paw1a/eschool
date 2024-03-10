include .env
export

build:
	go mod download && go build -o ./.bin/app ./cmd/app/main.go

run: build
	./.bin/app

.DEFAULT_GOAL := run
