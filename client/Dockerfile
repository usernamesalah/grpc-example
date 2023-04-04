FROM golang:alpine

WORKDIR /client

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o client .

CMD ["./client"]
