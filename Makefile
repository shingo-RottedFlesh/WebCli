

up:
	docker compose up --build



clean:
	docker system prune -a

remove:
	docker rm -f $(docker ps -a -q)
	docker compose rm -f $(docker compose ps -a -q)