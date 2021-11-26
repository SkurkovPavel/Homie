.PHONY: build
build:
	go build -v ./cmd/homie

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: migrate-up
migrate:
	migrate -path ./migrations -database "" up

.DEFAULT_GOAL := build