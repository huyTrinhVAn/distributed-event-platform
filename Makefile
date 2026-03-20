COMPOSE_FILE=deployments/docker-compose.yml

compose-up:
	docker compose -f $(COMPOSE_FILE) up -d

compose-down:
	docker compose -f $(COMPOSE_FILE) down

compose-logs:
	docker compose -f $(COMPOSE_FILE) logs -f

compose-ps:
	docker compose -f $(COMPOSE_FILE) ps

compose-restart:
	docker compose -f $(COMPOSE_FILE) restart

kafka-shell:
	docker exec -it kafka sh

volume-ls:
	docker volume ls
