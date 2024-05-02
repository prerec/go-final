FROM golang:1.22.1-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add gcc musl-dev

COPY ["go.mod", "go.sum", "./"]

RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . ./
RUN go build -o ./bin/todo_list ./cmd/main.go

FROM alpine AS runner

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /usr/local/src/bin/todo_list /bin
COPY ./configs /configs
COPY ./migrations /migrations
COPY ./web /web

RUN goose -dir ./migrations sqlite3 ./scheduler.db up

CMD ["./bin/todo_list"]