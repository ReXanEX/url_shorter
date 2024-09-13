# Этап сборки
FROM golang:1.23.1-alpine3.20 AS builder
WORKDIR /app
# Копируем go.mod и go.sum для кэширования зависимостей
COPY src/go.mod src/go.sum ./
# Загружаем зависимости
RUN go mod download
# Копируем исходный код
COPY src/ .

# Собираем приложение
RUN go build -o myapp main.go


# Этап выполнения
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/myapp .
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]

EXPOSE 8080
