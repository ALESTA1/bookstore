FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY server/ ./server
COPY db/ ./db
COPY proto/ ./proto

WORKDIR /app/server

RUN go build -o server .

CMD ["./server"]
