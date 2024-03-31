include .env
export

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" -o ./.bin/app ./cmd/app/main.go

run: build
	docker-compose up postgres redis app

debug: build
	docker-compose up postgres redis debug

migrate:
	# if "error: file does not exist" was occurred,
    # it means that data is up to date
	docker-compose up migrate

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


clean:
	rm -rf .bin .data

.DEFAULT_GOAL := run
