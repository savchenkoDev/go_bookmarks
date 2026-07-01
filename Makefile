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