.SILENT:

build: mod verify
	go build -o .bin/main ./cmd/app

mod:
	go mod tidy -v

verify:
	go mod verify

run: compose-build
	docker compose  up

stop:
	docker compose  down

restart:
	docker compose  restart

compose-build: build
	docker compose  build
