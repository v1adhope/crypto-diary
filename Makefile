.SILENT:

run: build
	./.bin/main

build: mod verify
	go build -o .bin/main ./cmd/app
mod:
	go mod tidy -v

verify:
	go mod verify

# Test docker network
n-run:
	docker compose up -d

n-stop:
	docker compose down

n-restart:
	docker compose restart
