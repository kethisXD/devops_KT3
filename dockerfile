# 1) Сборка статического бинаря
FROM golang:1.24.2 AS builder
WORKDIR /app



COPY . .

# Собираем под Linux, отключаем CGO → получаем полностью статический exe
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -ldflags="-s -w" -o web ./cmd/web

# 2) Минимальный финальный образ
FROM alpine:latest
# Если ваше приложение делает HTTPS-запросы, раскомментируйте:
# RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/web .

EXPOSE 8080
ENTRYPOINT ["./web"]
