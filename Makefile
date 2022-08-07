postgres:
	docker run --name postgresdb2 -p 5000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine3.16
createdb:
	docker exec -it postgresdb createdb --username=root --owner=root persons

dropdb:
	docker exec -it postgresdbc dropdb persons

migrateup:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5000/persons?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5000/persons?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY:  postgres createdb dropdb migrateup migratedown sqlc