build:
	@go build -o bin/go_ecommerce cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go_ecommerce