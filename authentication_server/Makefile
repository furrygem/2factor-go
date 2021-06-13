.PHONY: build
build:
	go build -v cmd/app/main.go
.PHONY: tests
TESTING_PASSWORD=`cat internal/app/store/testing_password.txt`
tests: tests_migrations_up
	go test -v ./...
tests_start_docker:
	docker run --rm --name postgres_testing_db -e POSTGRES_PASSWORD=$(TESTING_PASSWORD) -e POSTGRES_DB=tests -e POSTGRES_USER=postgres -p127.0.0.1:15432:5432 -d postgres
	
tests_migrations_down:
	migrate -path test_migrations -database postgres://postgres:$(TESTING_PASSWORD)@127.0.0.1:15432/tests?sslmode=disable down
tests_migrations_up:
	migrate -path test_migrations -database postgres://postgres:$(TESTING_PASSWORD)@127.0.0.1:15432/tests?sslmode=disable up
clean_tests: tests_migrations_down
	docker stop postgres_testing_db
DEFAULT_GOAL := build