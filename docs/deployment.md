# 設定・デプロイ

## 設定ファイル（yayp.toml）

TOML形式。デフォルトパスは `./yayp.toml`。`-config` フラグで変更可能。

```toml
[server]
    YPPrefix = "ap"           # YPプレフィックス（ジャンル解析に使用）
    Port     = "8000"         # HTTPサーバーのリッスンポート
    LogPath  = "peercast-yayp.log"  # ログファイルパス
    Debug    = true           # デバッグモード（詳細ログ出力）

[database]
    Host     = "localhost"    # MySQLホスト
    Port     = "3306"         # MySQLポート
    User     = "root"         # MySQLユーザー
    Password = "password"     # MySQLパスワード
    DB       = "yayp"         # データベース名

[peercast]
    Host         = "localhost"  # PeerCastデーモンのホスト
    Port         = "7144"       # PeerCastデーモンのポート
    AuthType     = "basic"      # 認証タイプ（"basic" or 空）
    AuthUser     = "root"       # Basic認証ユーザー
    AuthPassword = "peer"       # Basic認証パスワード
```

---

## コマンドライン引数

```
./yayp [-config <設定ファイルパス>]
```

| フラグ | デフォルト | 説明 |
|---|---|---|
| `-config` | `./yayp.toml` | 設定ファイルのパス |

---

## Docker によるデプロイ

### ビルド

```bash
docker build -t peercast-yayp .
```

マルチステージビルド：
1. `golang:1.25-alpine3.21` でバイナリをビルド
2. `alpine:3.21` に最小限のファイルだけコピー

### docker-compose

```bash
cd docker-compose
docker-compose up -d
```

#### サービス構成

| サービス | イメージ | ポート | 説明 |
|---|---|---|---|
| `yayp` | ビルドイメージ | 8000 | アプリケーション本体 |
| `mariadb` | mariadb:10.3 | 3306 | データベース |
| `peercast` | titagaki/peercast-yt | 7144 | PeerCastデーモン |

#### 環境変数（.env）

```
VOLUMES_PATH=<ホスト側のボリュームパス>
MYSQL_ROOT_PASSWORD=<MySQLrootパスワード>
MYSQL_DATABASE=<データベース名>
PEERCAST_PORT=<ホスト側PeerCastポート>
PEERCAST_PASSWORD=<PeerCastパスワード>
```

### DBスキーマ初期化

```bash
mysql -u root -p yayp < docker-compose/sql/schema.sql
mysql -u root -p yayp < docker-compose/sql/seed.sql
```

---

## Graceful Shutdown

1. `SIGINT` または `SIGTERM` を受信
2. Worker の `context.Context` をキャンセル（バックグラウンドジョブ停止）
3. HTTPサーバーを10秒タイムアウトでシャットダウン
4. プロセス終了

---

## フロントエンド

- git submodule: `frontend/` → `git@github.com:titagaki/peercast-yayp-frontend.git`
- ビルド済みファイルは `public/` に配置
- Echo の Static ミドルウェアで配信（HTML5 History Mode 対応）
- フロントエンドの更新手順：
  ```bash
  cd frontend
  npm install && npm run build
  cp -r dist/* ../public/
  ```
