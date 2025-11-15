binary = helitask

release:
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(binary)_windows_amd64
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(binary)_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(binary)_darwin_amd64

run:
	docker-compose up

dev:
	docker-compose -f docker-compose.yml up --build

prod:
	docker-compose -f docker-compose.prod.yml up --build

test:
	go test ./...

migrate:
	go run ./cmd/migrate -env=development

