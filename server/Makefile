
run-dev:
	go run ./cmd/

run: 
	./app

build: 
	go build -o app ./cmd/* 

test:
	go test ./internals/UserService... -v --race

coverage:
	go test ./internals/UserService... -v \
	-coverpkg=./internals/UserService/... \
	-coverprofile=coverage.out \
	./internals/UserService/...

coverage-details:
	go tool cover -func=coverage.out

coverage-html:
	go tool cover -html=coverage.out

db:
	docker-compose up -d

generate:
	sqlc generate

install-dependencies:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

