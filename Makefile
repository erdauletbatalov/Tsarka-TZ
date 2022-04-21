build:
	docker-compose build
	docker image prune

run:
	docker-compose build
	docker image prune
	docker-compose up

stop:
	docker-compose stop

delete:
	docker-compose down
	docker rmi redis:6.2-alpine postgres:latest alpine:latest texzadanie_app:latest