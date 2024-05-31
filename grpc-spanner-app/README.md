# Overview

gRPC / Cloud Spanner

# 動作確認

```bash
go run cmd/grpc/main.go
```

```bash
grpcurl -plaintext -import-path ./proto -proto service.proto -d '{"user_id": "12345"}' localhost:50051 service.UserService/GetUser
```

レスポンス例

```json
{
  "userId": "12345",
  "name": "John Doe"
}
```
