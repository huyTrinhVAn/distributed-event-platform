compose-up:
	docker compose -f deployments/docker-compose.yml up -d

compose-down:
	docker compose -f deployments/docker-compose.yml down

compose-logs:
	docker compose -f deployments/docker-compose.yml logs -f

kafka-shell:
	docker exec -it kafka sh