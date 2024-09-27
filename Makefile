include docker/docker.mk

base-up:
	docker network create --driver bridge development
base-down:
	docker network rm development

dev-up:
	docker compose -f dev/docker-compose.yaml up -d --build
dev-down:
	docker compose -f dev/docker-compose.yaml down

up: base-up services-up tools-up dev-up
down: dev-down tools-down services-down base-down