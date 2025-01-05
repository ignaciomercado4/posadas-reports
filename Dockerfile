FROM golang:1.23.3-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add postgresql-client

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/templates ./templates
COPY --from=builder /app/public ./public

EXPOSE 8080

CMD ["./main"]