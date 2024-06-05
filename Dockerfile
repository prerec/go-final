FROM golang:1.22.1-alpine AS builder

WORKDIR /app

RUN apk --no-cache add gcc musl-dev

COPY . .

RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . ./
RUN go build -o ./bin/todo_list ./cmd/main.go

FROM alpine AS runner

WORKDIR /app

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/bin/todo_list /app
COPY ./configs /app/configs
COPY ./migrations /app/migrations
COPY ./web /app/web

RUN goose -dir ./migrations sqlite3 ./scheduler.db up

CMD ["./todo_list"]