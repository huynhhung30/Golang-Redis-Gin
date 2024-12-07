run:
	go run cmd/main.go

up:
	docker-compose up -d 

down:
	docker-compose down

build:
	docker-compose build --no-cache
swag:
	swag init -g ./cmd/main.go -o cmd/docs