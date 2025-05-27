FROM golang:1.22.2 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y git gcc libc6-dev

ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=1

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o health-checker ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/health-checker .

CMD ["./health-checker"]
