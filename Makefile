dev-up:
	@docker compose \
		-f docker/docker-compose.yml up -d

dev-down:
	@docker compose \
		-f docker/docker-compose.yml down

migrate-up:
	@migrate -path migration -database "mysql://root:password@tcp(0.0.0.0:3306)/shoe" -verbose up

migrate-down:
	@migrate -path migration -database "mysql://root:password@(0.0.0.0:3306)/shoe" -verbose down

dev-run:
	@go run src/cmd/bot/main.go
