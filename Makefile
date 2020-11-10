dc-build:
	docker build -t among-go .
dc-run:
	docker run -it --rm --name among-go-run among-go
dc-run-d:
	docker run -d -it --rm --name among-go-run among-go
build:
	go build
run:
	go run main.go
env:
	cp .env.example .env