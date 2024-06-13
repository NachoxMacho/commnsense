run: build
	@./bin/dreampicai

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

hot: build
	@air

css:
	@tailwindcss -i view/css/styles.css -o public/styles.css --watch

templ:
	@templ generate --watch

build:
	@templ generate view
	@go build -tags dev -o bin/dreampicai main.go

up: ## Database migration up
	@go run cmd/migrate/main.go up

reset:
	@go run cmd/reset/main.go up

down: ## Database migration down
	@go run cmd/migrate/main.go down

migration: ## Migrations against the database
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

seed:
	@go run cmd/seed/main.go

