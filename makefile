postgres:
	docker run --name postgres18 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -v pgdata:/var/lib/postgresql/data -d postgres:18-alpine

create_db:
	docker exec -it postgres18 createdb --username=root --owner=root simple_bank

drop_db:
	docker exec -it postgres18 dropdb simple_bank

migrate_up:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate


.PHONY: postgres create_db drop_db migrate_up migrate_down sqlc