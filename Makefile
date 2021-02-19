build:
	go build cmd/api/*.go

run:
	docker-compose up -d
	go build cmd/api/*.go
	go run cmd/api/main.go

test:
	go test -v ./...
