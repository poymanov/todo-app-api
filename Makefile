.DEFAULT_GOAL := restart

init: docker-down-clear docker-build docker-up
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
	docker compose run --rm app su user

start-app:
	docker compose run --rm app go run cmd/main.go
