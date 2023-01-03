.SILENT:

# Test docker network
n-run:
	docker compose up -d

n-stop:
	docker compose down

n-restart:
	docker compose restart
