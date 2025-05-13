FROM golang:1.22.1-alpine AS builder

WORKDIR /app

# Install necessary tools
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .


# Build the Go app
RUN go build -o main .

# --- Stage 2: Run ---
FROM alpine:latest
    
WORKDIR /root/
    
# Copy the built binary from the builder stage
COPY --from=builder /app/main .
    
COPY --from=builder /app/.env ./
    
COPY --from=builder /app/migrations ./migrations
# Expose the port the app runs on (default Gin port)
EXPOSE 7170

# Run the Go app
WORKDIR /root/
CMD ["./main"]
