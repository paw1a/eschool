include .env
export

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run: build
	docker-compose up app redis postgres

migrate:
	# if "error: file does not exist" was occurred,
    # it means that data is up to date
	docker-compose up migrate

clean:
	rm -rf .bin .data

.DEFAULT_GOAL := run
