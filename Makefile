postgres:
	docker run --name postgres -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owne=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose up  

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  simple_bank/db/sqlc Store

.PHONY: createdb dropdb postgres migrateup migreatedown sqlc test server mock
#
