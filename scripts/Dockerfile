FROM golang:1.22.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/ein.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
