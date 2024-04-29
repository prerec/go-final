build:
	go build -o ./bin/todo_list ./cmd/main.go
run:
	goose -dir ./migrations sqlite3 ./scheduler.db up
	./bin/todo_list