.PHONY: default run build test docs clean
# Variables
APP_NAME=gopportunities

# Tasks
default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@if exist $(APP_NAME) del /F /Q $(APP_NAME)
	@if exist docs rmdir /S /Q docs
