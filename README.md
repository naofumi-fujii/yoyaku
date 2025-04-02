# 予約システム

カレンダーの一部を範囲選択して予約できるシステムです。

## 技術スタック

### バックエンド
- Go 1.24
- Echo (Webフレームワーク)
- MySQL 5.7
- Docker

### フロントエンド
- Next.js 14
- TypeScript 5
- Tailwind CSS 3
- FullCalendar (カレンダーコンポーネント)
- Storybook (UIコンポーネント開発・テスト)

## セットアップと実行方法

### Docker を使う場合（推奨）

```bash
# プロジェクトのルートディレクトリで実行
docker-compose up --build
```

これにより以下のサービスが起動します：
- フロントエンド: http://localhost:3000
- バックエンド: http://localhost:8080
- MySQL: localhost:3306

### 直接実行する場合

#### バックエンドのセットアップ

```bash
# バックエンドディレクトリに移動
cd backend

# 依存関係のインストール
go mod tidy

# サーバーの起動
go run cmd/server/main.go
```

バックエンドサーバーは http://localhost:8080 で起動します。

#### フロントエンドのセットアップ

```bash
# フロントエンドディレクトリに移動
cd frontend

# 依存関係のインストール
npm install

# 開発サーバーの起動
npm run dev

# Storybookの起動（UIコンポーネント開発・テスト用）
npm run storybook
```

フロントエンドは http://localhost:3000 で起動します。
Storybookは http://localhost:6006 で起動します。

## 機能

- トップページにカレンダーが表示されます
- カレンダーの一部を範囲選択すると予約ダイアログが表示されます
- 予約を確定すると、カレンダーに予約が表示されます
- 右側のパネルには予約の一覧が表示され、予約のキャンセルができます
- アトミックデザインパターンに基づいたコンポーネント設計
- Storybookによるコンポーネントのドキュメント化とテスト

## API

### 予約の作成
- エンドポイント: `POST /api/reservations`
- リクエスト:
  ```json
  {
    "startTime": "2023-04-01T10:00:00Z",
    "endTime": "2023-04-01T11:00:00Z"
  }
  ```

### 予約の一覧取得
- エンドポイント: `GET /api/reservations`

### 予約の削除
- エンドポイント: `DELETE /api/reservations/:id`

## 開発環境のデータベース設定

MySQL の接続情報:
- ホスト: localhost (Docker: mysql)
- ポート: 3306
- データベース名: reservations
- ユーザー名: root
- パスワード: password

## トラブルシューティング

### Docker 環境での一般的な問題解決

1. コンテナのログを確認
```bash
docker-compose logs -f
```

2. 特定のサービスのログを確認
```bash
docker-compose logs -f frontend
docker-compose logs -f backend
docker-compose logs -f mysql
```

3. コンテナを再ビルド
```bash
docker-compose down
docker-compose up --build
```

## プロジェクト構成

### フロントエンド
- アトミックデザインパターンを採用（詳細は `frontend/README-atomic-design.md` を参照）
- Storybookによるコンポーネント開発（詳細は `frontend/README-storybook.md` を参照）

### バックエンド
- クリーンアーキテクチャを採用
  - handler: APIエンドポイントの処理
  - service: ビジネスロジック
  - repository: データアクセス
  - model: データモデル

## テスト

### バックエンドのテストカバレッジ
```bash
# backendディレクトリに移動
cd backend

# テスト実行とカバレッジレポート生成
go test -v -coverprofile=coverage.out ./...

# カバレッジレポートを関数ごとに表示
go tool cover -func=coverage.out

# HTMLカバレッジレポート生成（オプション）
go tool cover -html=coverage.out -o coverage.html
```