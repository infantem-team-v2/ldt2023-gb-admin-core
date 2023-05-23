# Building sources first
FROM golang:1.20-alpine as BuildStage

WORKDIR /app

RUN ls -la
RUN apk add build-base

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Making build in other directory
RUN go build -o main ./cmd/api/main.go

# Running pre-build
FROM alpine:latest as RunStage

WORKDIR /app

COPY --from=BuildStage /app .

CMD ["./main"]
