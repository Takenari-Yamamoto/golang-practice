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

# Cloud Spanner

- インスタンス作成

```bash
gcloud spanner instances create spanner-sample-instance --config=regional-asia-northeast1 --description="初めての spanner" --nodes=1
```

- DB 作成

```bash
gcloud spanner databases create spanner-sample-db --instance=spanner-sample-instance
```

- テーブル作成

```bash
gcloud spanner databases ddl update spanner-sample-db --instance=spanner-sample-instance --ddl="$(cat schema/schema.sql)"
```

この場合は schema/\*\*.sql が配置されてる時に実行する

- 作成済のデータベース

```bash
gcloud spanner databases list --instance=spanner-sample-instance
```

参考: https://cloud.google.com/spanner/docs/create-query-database-console?hl=ja
