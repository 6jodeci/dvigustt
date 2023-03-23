include .env
export

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER="$(POSTGRES_USER)" -e POSTGRES_PASSWORD="$(POSTGRES_PASS)" -d postgres:14.4-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root "$(POSTGRES_NAME)"

dropdb:
	docker exec -it postgres14 dropdb "$(POSTGRES_NAME)"

docker:
	docker compose up

migratecreate:
	migrate create -ext sql -dir migrations -seq init_db

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

test:
	go test -v -cover ./...

swagger:
	swag init -g ./cmd/main.go -o ./docs

.PHONY: postgres createdb dropdb docker migratecreate migrateup migratedown test swagger