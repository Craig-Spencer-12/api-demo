getUsers:
	curl http://localhost:8080/users

user:
	curl -X POST http://localhost:8080/users \
	-H "Content-Type: application/json" \
	-d '{"username":"craigspencer", "email":"test@gmail.com"}'

consumer:
	go run cmd/consumer/main.go

server:
	go run cmd/api/main.go

docker:
	docker compose up -d