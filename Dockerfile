# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /src

# Install build tools
RUN apk add --no-cache ca-certificates git

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o helitask ./

# Final stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /src/helitask ./helitask

# Use non-root user for runtime security
RUN adduser -D -g '' helitask && chown helitask:helitask /app/helitask
USER helitask

EXPOSE 8080

ENV APP_ENV=production

ENTRYPOINT ["/app/helitask"]
