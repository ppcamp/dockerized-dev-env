

down:
	docker compose down

up: down
	docker compose up -d

upx: down
	docker compose up --build -d