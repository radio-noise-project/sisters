include .env

ci:
	golangci-lint run ./...
test:
	go test -v  ./...
coverage:
	go test -v -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
run:
	go run cmd/sisters/main.go 
clean:
	rm cover.html cover.out 
.PHONY: ci test coverage run
