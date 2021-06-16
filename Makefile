.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./app ./cmd/app/main.go

run:
	docker-compose up --remove-orphans app
