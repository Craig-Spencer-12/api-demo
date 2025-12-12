# ========== BUILD ===============
FROM golang:latest AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o consumer ./cmd/consumer


# ============ Runtime =============
FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/consumer /app/consumer

EXPOSE 9093

CMD ["./consumer"]
