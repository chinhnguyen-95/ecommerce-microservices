# Stage 1: Build
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum from the root directory of the repository
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project (including the product-service directory and common module)
COPY . .

# Change working directory to the product-service
WORKDIR /app/product-service

# Build the product-service binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o product-service .

# Stage 2: Run
FROM alpine:latest

# Install minimal dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/product-service/product-service .

# Expose the application port
EXPOSE 8001

# Command to run the application
CMD ["./product-service"]