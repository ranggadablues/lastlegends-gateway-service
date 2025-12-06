# ==========================
# 1. BUILD STAGE
# ==========================
FROM golang:1.24.7-alpine AS builder

# Install git (needed for go mod download)
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first (cache layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway-service ./cmd

# ==========================
# 2. RUN STAGE
# ==========================
FROM alpine:3:23.0

WORKDIR /app

# Copy binary from build stage
COPY --from=builder /app/gateway-service .

# Expose your gateway port (example: 8080)
EXPOSE 8080

# Run binary
CMD ["./gateway-service"]
