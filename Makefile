build:
	go build -o ./bin/todo_list ./cmd/main.go
run:
	goose -dir ./migrations sqlite3 ./scheduler.db up
	./bin/todo_list
docker_build:
	docker build -t todo-service .
docker_run:
	docker run -d -p 7540:7540 todo-service
goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest
test:
	go test ./tests