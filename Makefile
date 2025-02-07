.DEFAULT_GOAL := restart

init: docker-down-clear docker-build copy-configs docker-up migrate
up: docker-up
down: docker-down
restart: down up

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down --remove-orphans

docker-down-clear:
	docker-compose down -v --remove-orphans

docker-build:
	docker-compose build

shell:
	docker-compose exec todo bash

copy-configs:
	cp config/config.example.yml config/config.yml

migrate:
	docker-compose exec todo go run cmd/migrator/migrator.go

logs:
	docker-compose logs -f

generate-swagger:
	docker-compose exec todo swag i -g internal/app/app.go

format-swagger:
	docker-compose exec todo swag fmt

test:
	docker-compose exec todo go test -v ./...

lint:
	docker-compose exec todo golangci-lint run

check: lint test