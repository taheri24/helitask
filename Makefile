run:
	docker-compose up

dev:
	docker-compose -f docker-compose.yml up --build

prod:
	docker-compose -f docker-compose.prod.yml up --build

test:
	go test ./...

migrate:
	go run ./cmd/main.go -migrate=true -env=development