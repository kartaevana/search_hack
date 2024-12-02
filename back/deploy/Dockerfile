# Этап 1: Сборка Go-приложения
FROM golang:latest AS builder

WORKDIR /back

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./cmd/main ./cmd/main.go && chmod +x ./cmd/main

# Этап 2: Финальный образ
FROM alpine:latest

WORKDIR /back/cmd

# Копирование из этапа builder
COPY --from=builder /back/cmd/main ./main

EXPOSE 8080

CMD ["./main"]
