# Build stage
FROM golang:1.22.0 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN chmod +x wait-for-it.sh
RUN go build -o /go-exec

EXPOSE 8080

CMD ["./wait-for-it.sh", "mysql:3306", "--", "./wait-for-it.sh", "redis:6379", "--", "/go-exec"]
