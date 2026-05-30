.PHONY: help run build tidy docker-up docker-down docker-logs db-only clean

help:
	@echo "Targets:"
	@echo "  run          - Run app locally (requires .env and running Postgres)"
	@echo "  build        - Build binary to bin/server"
	@echo "  tidy         - go mod tidy"
	@echo "  docker-up    - Start Postgres + app with Docker Compose"
	@echo "  docker-down  - Stop Docker Compose services"
	@echo "  docker-logs  - Follow app container logs"
	@echo "  db-only      - Start only Postgres (for local dev)"
	@echo "  clean        - Remove build artifacts"

run:
	go run ./cmd/server

build:
	@mkdir -p bin
	go build -o bin/server ./cmd/server

tidy:
	go mod tidy

docker-up:
	docker compose up --build -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f app

db-only:
	docker compose up -d db

clean:
	rm -rf bin/
