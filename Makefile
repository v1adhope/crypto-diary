.SILENT:

default: build

mod:
	go mod tidy -v

verify:
	go mod verify

build: mod verify
	go build -o .bin/main -race ./cmd/app

stop:
	docker compose down -v

run: stop build
	docker compose up --build

# .PHONY: cover
# cover:
# 	go test -short -count=1 -race -coverprofile=cover.out ./...
# 	go tool cover -html=cover.out
# 	rm cover.out
