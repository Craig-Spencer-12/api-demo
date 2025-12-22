FROM golang:1.25.5-alpine

WORKDIR /app

COPY common/ /common/

# COPY go.mod go.sum ./

COPY example-service/ ./
RUN go mod download

RUN go build -o example-service-app ./cmd/main.go

CMD ["./example-service-app"]