.PHONY: build
build:
	go build -v ./cmd/rsoi-lab

.PHONY: migrate
migrate:
	goose -dir ./migrations postgres "user=program dbname=persons password=test host=localhost port=5432 sslmode=disable" up

.PHONY: test
test:
	go test -v ./...

.DEFAULT_GOAL := build
