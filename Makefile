.PHONY: dev build run test

build:
	@echo "Building go-shit"
	@go build -o build/go-shit cmd/go-shit/main.go
