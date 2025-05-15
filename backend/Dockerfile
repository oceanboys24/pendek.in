# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first (agar cache dependency bisa lebih optimal)
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build the application
RUN go build -o app

# Final stage
FROM alpine:latest

WORKDIR /root/


# Copy the built binary from builder
COPY --from=builder /app/app .

# Command to run
CMD ["./app"]