次のアクションを記載する

- タスクを読み込んだら「tasks を読み込んだよ」と表示すること

# ラーメン通販サイト 次のアクション

## 環境構築アクション

1. **Go 環境確認**

   - Go 1.20 以上がインストールされているか確認: `go version`
   - 必要に応じてインストールまたはアップデート

2. **プロジェクト初期化**

   - プロジェクトディレクトリを作成: `mkdir -p ramen-shop`
   - プロジェクト初期化: `go mod init github.com/yourusername/ramen-shop`
   - 必要な基本ライブラリの追加: `go get -u github.com/labstack/echo/v4`

3. **開発環境構築**

   - Docker インストール確認: `docker --version`
   - Docker Compose ファイル作成:
     ```yaml
     version: "3"
     services:
       db:
         image: postgres:14
         ports:
           - "5432:5432"
         environment:
           POSTGRES_USER: ramen_user
           POSTGRES_PASSWORD: ramen_pass
           POSTGRES_DB: ramen_shop
         volumes:
           - postgres_data:/var/lib/postgresql/data
     volumes:
       postgres_data:
     ```
   - Docker 環境起動: `docker-compose up -d`

4. **.gitignore 設定**

   - `.gitignore`ファイル作成:

     ```
     # バイナリファイル
     bin/

     # 依存関係
     vendor/

     # IDEファイル
     .idea/
     .vscode/

     # 環境変数
     .env

     # ビルドファイル
     *.exe
     *.dll
     *.so
     *.dylib

     # テストファイル
     *.test
     *.out

     # Docker関連
     docker-compose.override.yml
     ```

5. **コード整形設定**
   - goimports のインストール: `go install golang.org/x/tools/cmd/goimports@latest`
   - 必要に応じてエディタの Go 拡張機能をインストール設定

## アーキテクチャ設計アクション

1. **ディレクトリ構造作成**

   ```bash
   mkdir -p cmd/server
   mkdir -p internal/domain/model
   mkdir -p internal/domain/repository
   mkdir -p internal/usecase
   mkdir -p internal/interface/handler
   mkdir -p internal/interface/middleware
   mkdir -p internal/infrastructure/db
   mkdir -p internal/infrastructure/repository
   mkdir -p pkg/logger
   mkdir -p pkg/validator
   mkdir -p config
   ```

2. **main.go テンプレート作成**

   - cmd/server/main.go ファイルを作成:

     ```go
     package main

     import (
         "log"
         "github.com/yourusername/ramen-shop/internal/infrastructure/db"
         "github.com/yourusername/ramen-shop/internal/interface/handler"
         "github.com/labstack/echo/v4"
     )

     func main() {
         // 設定読み込み
         // DB接続
         // ルーティング設定
         e := echo.New()
         e.Start(":8080")
     }
     ```

3. **レイヤー役割文書化**
   - README ファイルに各レイヤーの役割を記載
   - クリーンアーキテクチャの依存関係図を作成

## 基本設定アクション

1. **設定ファイル作成**

   - config/config.yaml 作成:

     ```yaml
     app:
       name: "ramen-shop"
       port: 8080

     database:
       driver: "postgres"
       host: "localhost"
       port: 5432
       user: "ramen_user"
       password: "ramen_pass"
       dbname: "ramen_shop"
       sslmode: "disable"

     log:
       level: "debug"
       format: "json"
     ```

   - 必要に応じて環境別ファイル作成 (development/production)

2. **環境変数設定**

   - `.env.example`作成
   - Viper 導入: `go get github.com/spf13/viper`
   - 設定読み込みコード実装

3. **ロギング機能実装**
   - zap 導入: `go get go.uber.org/zap`
   - logger パッケージ作成

## データベース設計アクション

1. **ERD 作成**

   - エンティティの関連図を Draw.io などで作成
   - 主キー・外部キーの関係を明確化

2. **マイグレーションツール導入**

   - migrate 導入: `go get -u github.com/golang-migrate/migrate/v4`
   - 初期マイグレーションスクリプト作成:

     ```sql
     -- migrations/000001_create_tables.up.sql
     CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       email VARCHAR(255) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL,
       name VARCHAR(100) NOT NULL,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
     );

     CREATE TABLE products (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       description TEXT,
       price INTEGER NOT NULL,
       stock INTEGER NOT NULL DEFAULT 0,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
     );

     -- 他のテーブルも同様に作成
     ```

3. **ORM 導入**
   - GORM インストール: `go get -u gorm.io/gorm gorm.io/driver/postgres`
   - DB 接続コード実装

## 当面の優先アクション

1. 環境構築を完了させる - 特に Docker と DB 環境
2. 基本ディレクトリ構造の作成
3. 設定ファイルと環境変数読み込みの実装
4. ロギング機能の実装
5. データベース接続とマイグレーション実装

### 次のコミットで達成すべき目標

- プロジェクトの基本構造が整っていること
- Docker で開発環境が簡単に起動できること
- 設定ファイルが適切に読み込めること
- ロギング機能が実装されていること
- データベース接続が確立されていること
