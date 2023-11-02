build:
	go build -v ./cmd/app

migrate-up:
	migrate -path migrations -database "postgres://usr_admin:qwerty@localhost:5425/temp?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://usr_admin:qwerty@localhost:5425/temp?sslmode=disable" down

.PHONY: build migrate-up migrate-down
.DEFAULT_GOAL = build