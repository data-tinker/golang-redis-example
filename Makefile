build:
	docker build -t golang-redis-example .

run:
	docker-compose up --build

run-redis:
	docker pull redis & docker run -p 6379:6379 redis
