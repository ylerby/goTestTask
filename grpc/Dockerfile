FROM golang:1.21.1-alpine 
 
WORKDIR /cmd

COPY . .

RUN go mod download

WORKDIR ./cmd/grpc_server

RUN go build -o main . 
 
CMD ["./main"]

