.PHONY: dc-up dc-down

dc-up:
	docker compose -f devops/docker-compose.yml up -d

dc-down:
	docker compose -f devops/docker-compose.yml down
	docker volume rm devops_mongodata
