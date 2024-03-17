run-postgres:
	docker run --name postgres -e POSTGRES_DB=earthquake_db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres 

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/earthquake_db?sslmode=disable" -verbose up

create-schema:
	go get -d entgo.io/ent/cmd/ent
	go run -mod=mod entgo.io/ent/cmd/ent new
	go run -mod=mod ariga.io/entimport/cmd/entimport -h
	go run ariga.io/entimport/cmd/entimport -dsn "postgres://postgres:postgres@localhost:5432/earthquake_db?sslmode=disable"

generate-ent:
	go generate ./ent