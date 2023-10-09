# syntax=docker/dockerfile:1

# Build stage
# Notes: We want to deploy a small binary. This will keep our deployments as
#       fast as possible.

# Notes: Use this image to build the binary.
FROM golang:1.19-alpine3.17 AS builder
WORKDIR /app
COPY apps/lessons/dockerize .
RUN go build -o /app/dockerize

# Notes: Use the image below, which is smaller than the builder, to run our binary.
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app .

EXPOSE 8080
CMD [ "/app/dockerize", "--port", "8080" ]