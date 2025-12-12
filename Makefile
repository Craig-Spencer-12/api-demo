getUsers:
	curl http://localhost:8080/users

user:
	@read -p "Enter name: " USERNAME; \
	curl -X POST http://localhost:8080/users \
	-H "Content-Type: application/json" \
	-d '{"username":"'"$$USERNAME"'", "email":"'"$$USERNAME"'@gmail.com"}'

consumer:
	go run cmd/consumer/main.go

server:
	go run cmd/api/main.go

docker:
	docker compose up -d


peek_db: 
	docker compose exec -it postgres psql -U myuser -d mydb
# \dt - show all tables
