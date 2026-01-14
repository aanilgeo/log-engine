# Build the Go binary
FROM golang:1.21-alpine AS builder

# Install protoc dependencies if needed
RUN apk add --no-cache protobuf-dev

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the server and client
RUN go build -o server ./cmd/server/main.go
RUN go build -o client ./cmd/client/main.go

# Final lightweight image
FROM alpine:latest
WORKDIR /root/

# Copy binaries from the builder
COPY --from=builder /app/server .
COPY --from=builder /app/client .

# Expose gRPC port
EXPOSE 50051

# The command to run the server
CMD ["./server"]