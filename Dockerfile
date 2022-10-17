FROM golang:1.16-buster AS build

WORKDIR /app

RUN mkdir probagoyave
COPY ./probagoyave ./probagoyave
WORKDIR /app/probagoyave
RUN go mod download

RUN go build -o /docker-gs-ping

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-gs-ping /docker-gs-ping

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]
