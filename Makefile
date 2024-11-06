include .env

ci:
	golangci-lint run ./...
test:
	docker compose -f compose.test.yaml build --no-cache
	docker compose -f compose.test.yaml up
coverage:
	go test -v -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
run:
	docker compose -f compose.yaml build --no-cache
	docker compose -f compose.yaml up
clean:
	rm cover.html cover.out
	docker container prune -f
.PHONY: ci test coverage run
