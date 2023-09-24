FROM golang:1.21.1-alpine AS builder

RUN mkdir /planningapp

COPY . /planningapp

WORKDIR /planningapp

RUN CGO_ENABLED=0 go build -o planningservice ./cmd/main.go

RUN chmod +x /planningapp/planningservice

FROM alpine:latest

RUN mkdir /planningapp

COPY --from=builder /planningapp/planningservice /planningapp
COPY planning.env ./planning.env

EXPOSE 3000:3000

CMD ["/planningapp/planningservice"]