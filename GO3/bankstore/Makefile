pg_start:
	docker-compose start

pg_stop:
	docker-compose stop

createdb:
	docker-compose exec db createdb --username=postgres --owner=postgres bankstoredb

dropdb:
	docker-compose exec db dropdb -U postgres bankstoredb

migrateup:
	migrate -path ./db/migrations -database "postgresql://postgres:postgres@localhost:5433/bankstoredb?sslmode=disable" up

migratedown:
	migrate -path ./db/migrations -database "postgresql://postgres:postgres@localhost:5433/bankstoredb?sslmode=disable" down

.PHONY: pg_start pg_stop createdb dropdb migrateup migratedown