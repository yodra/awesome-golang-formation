build:
	go build cmd/api/*.go

run:
	go build cmd/api/*.go
	go run cmd/api/main.go

test:
	go test -v ./...
