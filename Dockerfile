FROM golang:1.20.2 AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o sesmate ./cmd/sesmate/

FROM alpine

COPY --from=builder /app/sesmate /usr/local/bin/sesmate

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

CMD ["sesmate"]
