
build:
	@go build -o bin/blog cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/blog

run:
	go run cmd/main.go

up:
	docker-compose up --build -d

down:
	docker-compose down

down-v:
	docker-compose down -v # remove even the volume