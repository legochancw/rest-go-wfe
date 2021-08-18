FROM golang:1.17.0-alpine

RUN     mkdir -p /app
WORKDIR /app
COPY . .
RUN    go mod download
RUN    go build -o app

EXPOSE 8090
ENTRYPOINT ["./app"]