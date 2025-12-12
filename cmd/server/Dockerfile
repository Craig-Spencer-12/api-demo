# ========== BUILD ===============
FROM golang:latest AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server


# ============ Runtime =============
FROM alpine:3.20

WORKDIR /app

# Add certificates if you make TLS requests
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server

EXPOSE 8080

CMD ["./server"]
