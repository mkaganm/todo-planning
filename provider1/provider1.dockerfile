FROM golang:1.21.1-alpine AS builder

RUN mkdir /provider1app

COPY . /provider1app

WORKDIR /provider1app

RUN CGO_ENABLED=0 go build -o provider1service ./cmd/main.go

RUN chmod +x /provider1app/provider1service

FROM alpine:latest

RUN mkdir /provider1app

COPY --from=builder /provider1app/provider1service /provider1app
COPY provider1.env ./provider1.env

EXPOSE 3001:3001

CMD ["/provider1app/provider1service"]