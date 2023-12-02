build:
	docker build -t golang-redis-example .

run:
	docker-compose up

run-redis:
	docker pull redis & docker run -p 6379:6379 redis
