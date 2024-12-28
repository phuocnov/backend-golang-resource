postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1234 -d postgres 
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:1234@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:1234@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc-generate:
	 sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc-generate
