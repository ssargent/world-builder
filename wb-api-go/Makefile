.PHONY: build
build:
	go build -o ./bin/backend ./cmd/backend

.PHONY: run-local
run-local: build
	./bin/backend  -e ./.env/development.env run

.PHONY: update-protos
update-protos:
	go get -u github.com/ssargent/apis

.PHONY: revendor
revendor: update-protos
	go mod tidy
	go mod vendor

.PHONY: lint
lint:
	golangci-lint run -c ./.golangci.yml