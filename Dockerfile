# Build stage
FROM golang:alpine AS builder

# Set working directory
WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Compile application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:latest

WORKDIR /root/

# Copy binary from build stage
COPY --from=builder /app/main .

# Copy configuration files
COPY --from=builder /app/config ./config

# Expose port
EXPOSE 8080

# Run application
CMD ["./main", "-e", "production"]
