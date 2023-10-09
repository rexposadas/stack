# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.19-alpine3.17 AS builder
WORKDIR /app
COPY apps/lessons/simple_service .
RUN go build -o /app/simple_service

FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app .

EXPOSE 8080
CMD [ "/app/simple_service", "--port", "8080" ]