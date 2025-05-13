FROM golang:1.22.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest
    
WORKDIR /root/

COPY --from=builder /app/main .
    
COPY --from=builder /app/.env ./

#  migrations
COPY --from=builder /app/migrations ./migrations

EXPOSE 7170

WORKDIR /root/
CMD ["./main"]
