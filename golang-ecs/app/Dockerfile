FROM golang:alpine

WORKDIR /app
COPY . .

RUN go mod download

WORKDIR /app/cmd
RUN go build -o main .

CMD ["/app/cmd/main"]
