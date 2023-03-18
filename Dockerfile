FROM golang:1.18-alpine AS builder
WORKDIR /build
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/build/discord-spam-bot main.go

FROM alpine:3.16
WORKDIR /app

COPY --from=builder /tmp/build/discord-spam-bot .
COPY .prod.env .
ENTRYPOINT ENV=prod /app/discord-spam-bot
