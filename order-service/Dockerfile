# Stage 1: Build
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum from the root directory of the repository
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project (including the order-service directory and common module)
COPY . .

# Change working directory to the order-service
WORKDIR /app/order-service

# Build the order-service binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o order-service .

# Stage 2: Run
FROM alpine:latest

# Install minimal dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/order-service/order-service .

# Expose the application port
EXPOSE 8003

# Command to run the application
CMD ["./order-service"]
