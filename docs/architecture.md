# peercast-yayp アーキテクチャ概要

## プロジェクト概要

PeerCast YP（Yellow Pages）サーバー。PeerCastネットワーク上で配信されているチャンネルの情報を収集し、クライアントやWebフロントエンドに提供する。

### 主な機能

- PeerCastデーモンから定期的にチャンネル情報をXMLで取得・同期
- チャンネル情報のインメモリキャッシュとDBへの永続化
- PeerCast互換 `index.txt` の生成・配信
- チャンネル一覧・ログのJSON API
- チャンネルログの10分間隔での記録
- 日次集計バッチ（最大リスナー数・平均リスナー数）
- Vue.js SPA フロントエンドのホスティング

---

## ディレクトリ構成

```
peercast-yayp/
├── cmd/
│   └── yayp/
│       └── main.go              # エントリポイント・DI ワイヤリング
├── internal/
│   ├── config/                  # 設定ファイルの読み込み
│   │   └── config.go
│   ├── domain/                  # ドメインエンティティ（ビジネスモデル）
│   │   ├── channel.go
│   │   ├── channel_log.go
│   │   ├── summary.go
│   │   └── information.go
│   ├── store/                   # データベース永続化層
│   │   ├── db.go
│   │   ├── channel.go
│   │   ├── channel_log.go
│   │   ├── summary.go
│   │   └── information.go
│   ├── cache/                   # インメモリキャッシュ層
│   │   ├── cache.go
│   │   ├── channel.go
│   │   └── information.go
│   ├── peercast/                # PeerCast外部連携クライアント
│   │   ├── client.go
│   │   ├── types.go
│   │   └── genre.go
│   ├── handler/                 # HTTPハンドラー
│   │   ├── handler.go
│   │   ├── channel.go
│   │   ├── channel_log.go
│   │   └── index_txt.go
│   ├── server/                  # HTTPサーバー構築・ルーティング
│   │   └── server.go
│   └── worker/                  # バックグラウンドジョブ
│       ├── worker.go
│       ├── sync_channel.go
│       └── daily_summary.go
├── public/                      # ビルド済みSPAフロントエンド
├── frontend/                    # フロントエンドのgit submodule
├── docker-compose/
│   ├── docker-compose.yml
│   └── sql/
│       ├── schema.sql
│       └── seed.sql
├── Dockerfile
├── go.mod
└── yayp.toml                   # 設定ファイル（サンプル）
```

---

## レイヤー構成

```
┌─────────────────────────────────────────┐
│              cmd/yayp/main.go           │  エントリポイント・DI
├─────────────────────────────────────────┤
│  server          │  worker              │  起動管理
├──────────────────┼──────────────────────┤
│  handler         │  (worker内ロジック)   │  ユースケース層
├──────────────────┴──────────────────────┤
│  store  │  cache  │  peercast (client)  │  インフラ層
├─────────┴─────────┴─────────────────────┤
│             domain                      │  ドメインエンティティ
├─────────────────────────────────────────┤
│             config                      │  設定
└─────────────────────────────────────────┘
```

### 依存関係の方向

- `cmd/yayp` → 全パッケージ（DIのワイヤリングのみ）
- `handler`, `worker` → `domain`（エンティティ参照）
- `handler`, `worker` → インターフェース経由で `store`, `cache`, `peercast` に依存
- `store`, `cache` → `domain`
- `domain` → 依存なし（最内層）

### インターフェース設計

Goの慣習に従い、**使用する側（consumer）がインターフェースを定義**：

| 定義場所 | インターフェース | 実装 |
|---|---|---|
| `handler` | `ChannelReader` | `store.ChannelStore` |
| `handler` | `ChannelLogReader` | `store.ChannelLogStore` |
| `handler` | `InformationReader` | `store.InformationStore` |
| `handler` | `ChannelCacheAccessor` | `cache.ChannelCache` |
| `handler` | `InformationCacheAccessor` | `cache.InformationCache` |
| `worker` | `ChannelRepository` | `store.ChannelStore` |
| `worker` | `ChannelLogRepository` | `store.ChannelLogStore` |
| `worker` | `SummaryRepository` | `store.SummaryStore` |
| `worker` | `ChannelCacheWriter` | `cache.ChannelCache` |
| `worker` | `PeercastFetcher` | `peercast.Client` |

---

## 依存性注入（DI）

`cmd/yayp/main.go` で全オブジェクトを手動組み立てし、各層に注入する。グローバル変数は使用しない。

```
main.go
 ├─ config.Load()           → *config.Config
 ├─ store.NewDB()           → *gorm.DB
 ├─ store.New*Store()       → 各Store
 ├─ cache.New()             → *gocache.Cache
 ├─ cache.New*Cache()       → 各Cache
 ├─ peercast.NewClient()    → *peercast.Client
 ├─ worker.New()            → *worker.Worker
 ├─ handler.New()           → *handler.Handler
 └─ server.New()            → *server.Server
```

---

## 技術スタック

| 領域 | ライブラリ |
|---|---|
| 言語 | Go 1.25 |
| HTTPフレームワーク | Echo v4 |
| ORM | GORM v2 |
| データベース | MySQL (MariaDB 10.3) |
| キャッシュ | go-cache（インメモリ） |
| 設定ファイル | TOML (BurntSushi/toml) |
| ロギング | log/slog（標準ライブラリ） |
| コンテナ | Docker multi-stage build (Alpine) |
