postgres:
	docker run --name postgresdb2 -p 5001:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine3.16
createdb:
	docker exec -it postgresdb2 createdb --username=root --owner=root persons

dropdb:
	docker exec -it postgresdb2 dropdb persons

migrateup:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5001/persons?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5001/persons?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY:  postgres createdb dropdb migrateup migratedown sqlc