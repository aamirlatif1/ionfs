build:
	@go build -o bin/ionfs

run: build
	@./bin/ionfs

test:
	@go test ./... -v