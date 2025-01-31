.DEFAULT_GOAL := restart

init: docker-down-clear docker-build copy-configs docker-up
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

logs:
	docker-compose logs -f