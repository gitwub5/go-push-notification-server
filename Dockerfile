# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy the Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application (this creates a static binary)
RUN go build -o push-server .

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /app/push-server .

# Expose the port the server runs on
EXPOSE 8080

# Command to run the server
CMD ["./push-server"]