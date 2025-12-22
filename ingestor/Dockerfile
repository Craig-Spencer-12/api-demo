FROM golang:1.25.5-alpine

WORKDIR /app

COPY common/ /common/

# COPY go.mod go.sum ./

COPY ingestor/ ./
RUN go mod download

RUN go build -o ingestor-app .

CMD ["./ingestor-app"]