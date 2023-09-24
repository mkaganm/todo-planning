FROM golang:1.21.1-alpine AS builder

RUN mkdir /provider2app

COPY . /provider2app

WORKDIR /provider2app

RUN CGO_ENABLED=0 go build -o provider2service ./cmd/main.go

RUN chmod +x /provider2app/provider2service

FROM alpine:latest

RUN mkdir /provider2app

COPY --from=builder /provider2app/provider2service /provider2app
COPY provider2.env ./provider2.env

EXPOSE 3002:3002

CMD ["/provider2app/provider2service"]