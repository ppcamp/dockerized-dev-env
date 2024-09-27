services-up:
	docker compose -f docker/services/docker-compose.yaml up -d
services-down:
	docker compose -f docker/services/docker-compose.yaml down

tools-up:
	docker compose -f docker/tools/docker-compose.yaml up -d
tools-down:
	docker compose -f docker/tools/docker-compose.yaml down