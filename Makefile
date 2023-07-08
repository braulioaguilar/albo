.PHONY: dc-up dc-down

dc-up:
	docker compose -f docker-compose.yml up -d

dc-down:
	docker compose -f docker-compose.yml down
	docker volume rm devops_mongodata
