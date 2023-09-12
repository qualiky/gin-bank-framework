postgres16:
	docker run --name psql16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Root_123 -d postgres:16beta2-alpine3.18

createdb:
	docker exec -it psql16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it psql16 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:Root_123@localhost:5432/simple_bank?sslmode=disable" -verbose up

sqlc:
	sqlc generate

migratedown:
	migrate -path db/migration -database "postgresql://root:Root_123@localhost:5432/simple_bank?sslmode=disable" -verbose down

test:
	go test -v -coverage ./...

server:
	go run main.go

.PHONY:
	postgres createdb dropdb migrateup migratedown sqlc server