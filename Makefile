include .env
export

up:
	docker compose -f ./docker-compose.yaml up -d --build

down:
	docker compose -f ./docker-compose.yaml down --remove-orphans

down-v:
	docker compose -f ./docker-compose.yaml down --remove-orphans --volumes

seed:
	(cd server && go run ./cmd/scripts/. -n $(RECORDS))
