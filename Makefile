.PHONY: migrate-up migrate-down

migrate-up:
	migrate -verbose -path ./schema -database $(DATABASE_URL) up

migrate-down:
	migrate -verbose -path ./schema -database $(DATABASE_URL) down
