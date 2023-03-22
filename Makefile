.SILENT:

build: mod verify
	go build -o .bin/main -race ./cmd/app

mod:
	go mod tidy -v

verify:
	go mod verify

run: compose-build
	docker compose  up

stop:
	docker compose  down

restart: stop run

compose-build: build
	docker compose  build

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=cover.out ./...
	go tool cover -html=cover.out
	rm cover.out
