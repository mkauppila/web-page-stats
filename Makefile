generate:
	go generate ./...

build:
	go build ./...

deps:
	go mod download

build-release:
	CGO_ENABLED=1 go build -ldflags="-s -w" -o ./bin/web-stats ./cmd/web-page-stats/main.go
		
run:
	CGO_ENABLED=1 go run cmd/web-page-stats/main.go

