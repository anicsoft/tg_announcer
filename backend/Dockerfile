# Build Stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev make

# Set up the working directory
WORKDIR /app

# Copy go.mod and go.sum files for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .
COPY Makefile ./

# Build the Go binary
RUN make build

# Final Stage
FROM alpine:latest

# Install necessary runtime dependencies (if any)
# Example: RUN apk add --no-cache curl

# Set up the working directory
WORKDIR /app

# Copy built binary from the builder stage
COPY --from=builder /app/build/application /app/

# Copy environment file and migrations directory from the builder stage
COPY --from=builder /app/.env /app/
COPY --from=builder /app/migrations /app/migrations

# Expose the port
EXPOSE 8888

# Command to run the application
CMD ["./application", "--host", "0.0.0.0", "--port", "8888"]