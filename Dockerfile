# Этап, на котором выполняется сборка приложения
FROM golang:1.16-alpine as builder
WORKDIR /build
# COPY ./probagoyave/go.mod . # go.sum
# RUN go mod download
COPY ./probagoyave .
RUN go build -o /main main.go
# Финальный этап, копируем собранное приложение
FROM alpine:3
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]
