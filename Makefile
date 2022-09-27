d.up:
	docker compose -f docker-compose.dev.yml up 

d.down:
	docker compose -f docker-compose.dev.yml down

d.up.build:
	docker compose -f docker-compose.dev.yml up --build