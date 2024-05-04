# client
FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o client .
CMD ["./client"]

# server
FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server .
CMD ["./server"]