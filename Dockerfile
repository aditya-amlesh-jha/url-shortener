FROM golang:1.22.0 AS builder

WORKDIR /app

COPY /server/* ./

RUN go mod download
RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]