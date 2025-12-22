FROM golang:1.25.5-alpine

WORKDIR /app

# COPY go.mod go.sum ./

COPY . .
RUN go mod download

RUN go build -o migration-app .

CMD ["./migration-app"]