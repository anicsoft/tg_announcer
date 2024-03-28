# Build Stage
FROM golang:alpine AS build

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set up the working directory
WORKDIR /app

# Copy only necessary files for building
COPY . .
COPY cmd/companies_service ./cmd/companies_service

# Build the Go binary
RUN CGO_ENABLED=1 CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go build -o ./build ./cmd/companies_service

# Final Stage
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk add --no-cache

# Set up the working directory
WORKDIR /app

# Copy built binary from the build stage
COPY --from=build /app/build/companies_service /app/

# Expose the port
EXPOSE 8080

# Command to run the application
CMD ["./companies_service"]