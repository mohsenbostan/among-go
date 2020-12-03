dc-build:
	docker-compose build
dc-run:
	docker-compose up
dc-run-d:
	docker-compose up -d
build:
	go build
run:
	go run main.go
env:
	cp .env.example .env