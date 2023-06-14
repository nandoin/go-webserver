FROM golang:1.16-alpine AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o webserver main.go

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/webserver .
CMD ["./webserver"]
