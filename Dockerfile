# Stage 1: Build the Go application
FROM golang:1.20.4-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -tags with_utls -o main ./start/main.go

# Stage 2: Create the final image using Alpine
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
