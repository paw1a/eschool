include .env
export

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" -o ./.bin/app ./cmd/app/main.go

run: build storage
	docker-compose up app

debug: build storage
	docker-compose up debug

storage:
	docker-compose up redis postgres

migrate:
	# if "error: file does not exist" was occurred,
    # it means that data is up to date
	docker-compose up migrate

clean:
	rm -rf .bin .data

.DEFAULT_GOAL := run
