test:
	docker compose run --rm test

migrate:
	docker compose run --rm db-migration

run: migrate
	go run ./cmd/api
