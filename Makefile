ifneq (,$(wildcard ./.env))
    include .env
    export
endif

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)


migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)?sslmode=disable" down 1

buildd:
	docker compose -f docker-compose.dev.yaml build

upd:
	docker compose -f docker-compose.dev.yaml up -d

downd:
	docker compose -f docker-compose.dev.yaml down

debug:
	docker compose -f docker-compose.dev.yaml logs -f app

shell:
	docker compose -f docker-compose.dev.yaml exec app sh

init:
	docker compose exec db psql -U postgres -d postgres -c "CREATE USER postgres WITH PASSWORD 'postgres';"