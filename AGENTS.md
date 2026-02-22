# AGENTS.md

## ビルド・テスト

```bash
# 依存解決
go mod tidy

# ビルド
go build ./cmd/yayp

# テスト（全パッケージ）
go test ./...

# テスト（詳細出力）
go test ./... -v
```

## プロジェクト構造

```
cmd/yayp/main.go        エントリポイント・DIワイヤリング
internal/config/        設定ファイルの読み込み（TOML）
internal/domain/        ドメインエンティティ（依存なし）
internal/store/         DB永続化層（GORM v2 + MySQL）
internal/cache/         インメモリキャッシュ層（go-cache）
internal/peercast/      PeerCastクライアント・ジャンルパーサー
internal/handler/       HTTPハンドラー（Echo v4）
internal/server/        サーバー構築・ルーティング
internal/worker/        バックグラウンドジョブ（同期・集計）
docs/                   アーキテクチャ・API・DB・デプロイ仕様
```

## 重要な規約

- インターフェースは**使用する側**（`handler`, `worker`）で定義する
- グローバル変数は使わない。依存関係は `main.go` でDIする
- スパゲッティは作らず、`internal/` 外への依存は `cmd/` のみ
- パッケージ名の typo に注意（旧 `repositoriy` は削除済み）
- ロガーは `log/slog`（標準ライブラリ）を使う
