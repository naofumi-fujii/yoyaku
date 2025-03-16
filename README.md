# 予約システム

カレンダーの一部を範囲選択して予約できるシステムです。

## 技術スタック

### バックエンド
- Go
- Echo (Webフレームワーク)

### フロントエンド
- Next.js
- TypeScript
- Tailwind CSS
- FullCalendar (カレンダーコンポーネント)

## セットアップと実行方法

### バックエンドのセットアップ

```bash
# バックエンドディレクトリに移動
cd backend

# 依存関係のインストール
go mod tidy

# サーバーの起動
go run cmd/server/main.go
```

バックエンドサーバーは http://localhost:8080 で起動します。

### フロントエンドのセットアップ

```bash
# フロントエンドディレクトリに移動
cd frontend

# 依存関係のインストール
npm install

# 開発サーバーの起動
npm run dev
```

フロントエンドは http://localhost:3000 で起動します。

## 機能

- トップページにカレンダーが表示されます
- カレンダーの一部を範囲選択すると予約ダイアログが表示されます
- 予約を確定すると、カレンダーに予約が表示されます
- 右側のパネルには予約の一覧が表示され、予約のキャンセルができます

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