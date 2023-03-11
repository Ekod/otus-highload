up:
	docker compose -f zarf/docker-compose.yaml up -d --build

migrate:
	migrate -path zarf/migrations -database "mysql://root:password@tcp(localhost)/social" up

run:
	go run cmd/app/main.go

down:
	docker compose -f zarf/docker-compose.yaml down --remove-orphans