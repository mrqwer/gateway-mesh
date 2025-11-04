# ==============================
# Stage 1: Build
# ==============================
FROM golang:1.22 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build based on target service
ARG SERVICE
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/service ./cmd/${SERVICE}

# ==============================
# Stage 2: Runtime
# ==============================
FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy binary
COPY --from=builder /bin/service /usr/local/bin/service

# Expose default port (can be overridden by compose)
EXPOSE 8080 50051 50052 50053 50054 50055

ENTRYPOINT ["/usr/local/bin/service"]
