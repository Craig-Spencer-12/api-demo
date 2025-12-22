FROM golang:1.25.5-alpine

WORKDIR /app

# COPY common/ /common/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o migration-app .

CMD ["./migration-app"]