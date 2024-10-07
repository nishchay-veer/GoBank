postgres:
	docker run --name postgres12 --network bank-network -e POSTGRES_PASSWORD=nish123 -e POSTGRES_USER=root -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migratedown:
	migrate -path db/migration -database "postgresql://root:nish123@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup:
	migrate -path db/migration -database "postgresql://root:nish123@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:nish123@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:nish123@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nishchay-veer/simplebank/db/sqlc Store

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--openapiv2_out=doc/swagger \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

evans:
	evans --host localhost --port 5000 -r repl --package pb --service SimpleBank

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine
	
.PHONY: postgres createdb dropdb sqlc test migrateup migratedown server mock proto migrateup1 migratedown1 proto evans