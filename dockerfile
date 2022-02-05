FROM golang:1.17.6-alpine3.15 as builder

WORKDIR /app

COPY . .

RUN go build .

FROM alpine:3.15 as runner

WORKDIR /app
COPY --from=builder /app/goping /app/goping