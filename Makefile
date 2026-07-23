include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d db

env-down:
	@docker compose down db

env-cleanup:
	@read -p "Remove all volume data? [y/N]?: " ans; \
	if [ "$$ans" = "y" ]; then \
	  docker compose down db && \
	  rm -rf out/pgdata && \
	  echo "Volume data removed"; \
	else \
	  echo "Canceled"; \
	fi

env-port-forwarder:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

# migrations
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "seq is required. Example: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm migrations \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "action is required. Example: make migrate-action action=\"down 2\" "; \
		exit 1; \
	fi; \
	docker compose run --rm migrations \
    		-path /migrations -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable \
    		"$(action)"

connect-db:
	@docker compose exec db \
	 psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}