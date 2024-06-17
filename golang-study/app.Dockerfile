# ----------------------------------------------------------------
# 開発用の Dockerfile はライブラリが入ってるのでデプロイする時はこれを使う
# ----------------------------------------------------------------

# build stage
FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

# final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
