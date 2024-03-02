install:
	@go mod download

build:
	@go build -o bin/slb ./cmd/slb/main.go

run-slb: build
	@./bin/slb
