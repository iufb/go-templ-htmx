build:
	@go build -o bin/go-templ-htmx cmd/main.go
test:
	@go test -v ./...
run: build
	@./bin/go-templ-htmx
