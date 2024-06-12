# Build Stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set up the working directory
WORKDIR /app

# Copy go.mod and go.sum files for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary
RUN go build -o ./build/ ./cmd/application

# Final Stage
FROM alpine:latest

# Set up the working directory
WORKDIR /app

# Copy built binary from the builder stage
COPY --from=builder /app/build/application /app/
COPY --from=builder /app/.env /app/

# Expose the backend port
EXPOSE ${BACKEND_PORT}

# Command to run the application
CMD ["./application", "--host", "0.0.0.0", "--port", "${BACKEND_PORT}"]
