FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod tidy && go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o app .

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache go ca-certificates

COPY --from=builder /app/app .

ENTRYPOINT ["/app/app"]
