include .env

.PHONY: run
run:
	docker compose -f compose.yaml build --no-cache
	docker compose -f compose.yaml up

.PHONY: stop
stop:
	docker compose -f compose.yaml down

.PHONY: run-development
run-development:
	docker compose -f compose.development.yaml build --no-cache
	docker compose -f compose.development.yaml up

.PHONY: stop-development
stop-development:
	docker compose -f compose.development.yaml down --rmi all --volumes --remove-orphans

.PHONY: protoc
protoc:
	docker build -t rnp/protobuf -f docker/development/protobuf/Dockerfile .
	docker run --rm -v $$PWD:/work -w /work/internal/api/proto rnp/protobuf \
	protoc --go_out=../runtime --go_opt=paths=source_relative \
	--go-grpc_out=../runtime --go-grpc_opt=paths=source_relative \
	runtime.proto

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	docker compose -f compose.test.yaml build --no-cache
	docker compose -f compose.test.yaml up
	
.PHONY: coverage
coverage:
	go test -v -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
	open cover.html

.PHONY: clean
clean:
	rm cover.html cover.out
	docker container prune -f
