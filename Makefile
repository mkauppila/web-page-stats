generate:
	go generate ./...
		
run:
	CGO_ENABLED=1 go run cmd/web-page-stats/main.go

