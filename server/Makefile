
run:
	go run ./cmd/

test:
	go test ./internals/UserService... --race

db:
	docker-compose up -d

generate:
	sqlc generate

sqlc-install:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

