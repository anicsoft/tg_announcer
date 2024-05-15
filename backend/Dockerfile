# Build Stage
FROM golang:1.19-alpine AS builder

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
RUN go build -o /app/build/companies_service ./cmd/companies_service

# Final Stage
FROM alpine:latest

# Install necessary runtime dependencies (if any)
# Example: RUN apk add --no-cache curl

# Set up the working directory
WORKDIR /app

# Copy built binary from the builder stage
COPY --from=builder /app/build/companies_service /app/

# Copy environment file and migrations directory from the builder stage
COPY --from=builder /app/.env /app/
COPY --from=builder /app/migrations /app/migrations

# Expose the port
EXPOSE 8888

# Command to run the application
CMD ["./companies_service", "--host", "0.0.0.0", "--port", "8888"]