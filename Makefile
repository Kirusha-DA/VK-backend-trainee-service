server:
	go run cmd/filmoteka/main.go
postgres-up:
	docker run --name=postgres --rm -e POSTGRES_USER=kirill -e POSTGRES_DB=filmoteka -e POSTGRES_PASSWORD=123 -v pgdata:/var/lib/postgresql/data -p 5432:5432 postgres:alpine 
postgres-down:
	docker stop postgres