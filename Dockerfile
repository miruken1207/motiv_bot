# Используем официальный образ Go для сборки
FROM golang:1.21-alpine AS builder

# Устанавливаем зависимости
RUN apk add --no-cache git

# Создаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum (если есть)
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o telegram-bot

# Финальный образ
FROM alpine:latest

# Устанавливаем зависимости для времени
RUN apk add --no-cache tzdata

# Копируем бинарный файл из builder
COPY --from=builder /app/telegram-bot /telegram-bot

# Указываем рабочую директорию
WORKDIR /

# Запускаем приложение
CMD ["/telegram-bot"]