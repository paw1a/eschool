include .env
export

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" -o ./.bin/app ./cmd/web/main.go

run: build
	docker-compose up postgres redis minio pgadmin app

debug: build
	docker-compose up postgres redis minio pgadmin debug

console: run_console

build_console:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" -o ./.bin/app ./cmd/console/main.go

run_console: build_console
	docker-compose up postgres redis minio pgadmin app

debug_console: build_console
	docker-compose up postgres redis minio pgadmin debug

migrate:
	# if "error: file does not exist" was occurred,
    # it means that data is up to date
	docker compose up migrate

mocks:
	mockery --dir internal/core/port --name IUserRepository --output internal/core/service/mocks \
		--filename user.go --structname UserRepository
	mockery --dir internal/core/port --name ICourseRepository --output internal/core/service/mocks \
		--filename course.go --structname CourseRepository
	mockery --dir internal/core/port --name ISchoolRepository --output internal/core/service/mocks \
		--filename school.go --structname SchoolRepository
	mockery --dir internal/core/port --name IReviewRepository --output internal/core/service/mocks \
		--filename review.go --structname ReviewRepository
	mockery --dir internal/core/port --name ILessonRepository --output internal/core/service/mocks \
		--filename lesson.go --structname LessonRepository
	mockery --dir internal/core/port --name IStatRepository --output internal/core/service/mocks \
		--filename stat.go --structname StatRepository
	mockery --dir internal/core/port --name IObjectStorage --output internal/core/service/mocks \
		--filename storage.go --structname ObjectStorage

clean:
	rm -rf .bin .data logs

.DEFAULT_GOAL := run