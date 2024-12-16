run: 
	go run main.go api
migrate:
	go run main.go migrate
lint:
	golangci-lint run
docker-up:
	docker compose up 
swagger:
	swag init