include .env
export

build_web:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" -o ./.bin/app ./cmd/web/main.go

run_web: build_web
	docker-compose up postgres redis minio pgadmin app

debug_web: build_web
	docker-compose up postgres redis minio pgadmin debug

console: run_console

build_console:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" -o ./.bin/app ./cmd/console/main.go

run_console: build_console
	docker-compose up postgres redis minio pgadmin app

debug_console: build_console
	docker-compose up postgres redis minio pgadmin debug

research:
	go mod download && go run cmd/research/main.go

migrate:
	# if "error: file does not exist" was occurred,
    # it means that data is up to date
	docker compose up migrate

mocks:
	mockery --dir internal/core/port --name IUserRepository --output internal/adapter/repository/mocks \
		--filename user.go --structname UserRepository
	mockery --dir internal/core/port --name ICourseRepository --output internal/adapter/repository/mocks \
		--filename course.go --structname CourseRepository
	mockery --dir internal/core/port --name ISchoolRepository --output internal/adapter/repository/mocks \
		--filename school.go --structname SchoolRepository
	mockery --dir internal/core/port --name IReviewRepository --output internal/adapter/repository/mocks \
		--filename review.go --structname ReviewRepository
	mockery --dir internal/core/port --name ILessonRepository --output internal/adapter/repository/mocks \
		--filename lesson.go --structname LessonRepository
	mockery --dir internal/core/port --name ICertificateRepository --output internal/adapter/repository/mocks \
		--filename certificate.go --structname CertificateRepository
	mockery --dir internal/core/port --name IObjectStorage --output internal/adapter/storage/mocks \
    		--filename storage.go --structname ObjectStorage
	mockery --dir internal/core/port --name IStatRepository --output internal/adapter/repository/mocks \
    		--filename stat.go --structname StatRepository

clean:
	rm -rf .bin .data logs

.DEFAULT_GOAL := run_web
