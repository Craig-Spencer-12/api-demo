sim:
	python3 simulation/bridge.py
.PHONY: sim

peek_db: 
	docker compose exec -it postgres psql -U myuser -d mydb
.PHONY: peek_db
# \dt - show all tables

peek_redis: 
	docker compose exec -it redis redis-cli
.PHONY: peek_redis
# GET key - SET key value - LRANGE key start stop

