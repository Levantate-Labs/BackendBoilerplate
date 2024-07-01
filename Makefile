build:
	@go build -o bin/server

run: build
	@./bin/server

run-db:
	@docker run --name server_db_instance -e POSTGRES_USER=server_user -e POSTGRES_PASSWORD=server_pass -e POSTGRES_DB=server_db -p 5432:5432 -d postgres
