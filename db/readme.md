Migration commands:

For up migration:

migrate -path db/migration -database "postgresql://root:Root_123@localhost:5432/postgres?sslmode=disable" -verbose up


For down migration:

migrate -path db/migration -database "postgresql://root:Root_123@localhost:5432/postgres?sslmode=disable" -verbose down

