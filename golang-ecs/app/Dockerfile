FROM --platform=linux/amd64 golang:alpine as builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64  

RUN go mod download
RUN go build -o main ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]
