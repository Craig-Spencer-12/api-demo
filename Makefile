getUsers:
	curl http://localhost:8080/users
.PHONY: getUsers

user:
	@USERNAME=$(word 2,$(MAKECMDGOALS)); \
	if [ -z "$$USERNAME" ]; then \
		read -p "Enter name: " USERNAME; \
	fi; \
	curl -X POST http://localhost:8080/users \
	-H "Content-Type: application/json" \
	-d '{"username":"'"$$USERNAME"'", "email":"'"$$USERNAME"'@gmail.com"}'
.PHONY: user

consumer:
	go run cmd/consumer/main.go
.PHONY: consumer

server:
	go run cmd/api/main.go
.PHONY: server

docker:
	docker compose up -d
.PHONY: docker


peek_db: 
	docker compose exec -it postgres psql -U myuser -d mydb
.PHONY: peek_db
# \dt - show all tables

%:
	@:
