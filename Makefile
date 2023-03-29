.SILENT:

build: mod verify
	go build -o .bin/main -race ./cmd/app

mod:
	go mod tidy -v

verify:
	go mod verify

run: stop build
	docker compose  up --build

stop:
	docker compose  down -v

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=cover.out ./...
	go tool cover -html=cover.out
	rm cover.out
