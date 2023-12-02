build:
	docker build -t golang-redis-example .

run:
	docker-compose up --build
