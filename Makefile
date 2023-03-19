include app.env
export

clickhouse:
	docker run -d --name dvigus-click -p 8123:8123 --ulimit nofile=262144:262144 yandex/clickhouse-server

make createdb:
	docker exec -it dvigus-click bash -c 'clickhouse-client --query "CREATE DATABASE IF NOT EXISTS dvigus_db"'

make dropdb:
	docker exec -it dvigus-click bash -c 'clickhouse-client --query "DROP DATABASE IF EXISTS dvigus_db"'

# Использование clickhouse драйвера для goose
migratecreate:
	goose -dir ./migrations clickhouse "$(DB_URL)" create ${name} sql

migrateup:
	goose -dir ./migrations clickhouse "$(DB_URL)" up

migratedown:
	goose -dir ./migrations clickhouse "$(DB_URL)" down


.PHONY: clickhouse createdb migratecreate migrateup migratedown
