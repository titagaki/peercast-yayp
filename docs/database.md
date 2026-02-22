# データベース設計

## 概要

- RDBMS: MySQL / MariaDB 10.3
- 文字セット: utf8
- ストレージエンジン: InnoDB
- ORM: GORM v2（テーブル名は複数形、snake_case）

---

## テーブル一覧

| テーブル名 | 説明 | 対応エンティティ |
|---|---|---|
| `channels` | 現在・過去のチャンネル情報 | `domain.Channel` |
| `channel_logs` | 10分間隔のチャンネルスナップショット | `domain.ChannelLog` |
| `channel_daily_summaries` | チャンネル日別集計 | `domain.ChannelDailySummary` |
| `information` | お知らせ情報（論理削除対応） | `domain.Information` |

---

## channels

チャンネルの最新状態を保持する。`is_playing` で配信中かどうかを管理。

| カラム | 型 | 説明 |
|---|---|---|
| `id` | INT UNSIGNED PK AI | 主キー |
| `cid` | VARCHAR(32) | PeerCastチャンネルID（32文字HEX） |
| `name` | VARCHAR(255) | チャンネル名（インデックス付き） |
| `bitrate` | INT | ビットレート (kbps) |
| `content_type` | VARCHAR(255) | コンテナタイプ (WMV, FLV, MKV等) |
| `listeners` | INT | リスナー数 |
| `relays` | INT | リレー数 |
| `age` | INT UNSIGNED | 配信時間（秒） |
| `genre` | VARCHAR(255) | ジャンル（YPプレフィックス除去後） |
| `description` | VARCHAR(255) | 概要 |
| `url` | VARCHAR(255) | コンタクトURL |
| `comment` | VARCHAR(255) | 配信者コメント |
| `track_artist` | VARCHAR(255) | トラック アーティスト |
| `track_title` | VARCHAR(255) | トラック タイトル |
| `track_album` | VARCHAR(255) | トラック アルバム |
| `track_genre` | VARCHAR(255) | トラック ジャンル |
| `track_contact` | VARCHAR(255) | トラック コンタクトURL |
| `hidden_listeners` | TINYINT(1) | リスナー数非表示フラグ |
| `tracker_ip` | VARCHAR(53) | トラッカーのIP:Port |
| `tracker_direct` | TINYINT(1) | 直接接続可否 |
| `is_playing` | TINYINT(1) | 配信中フラグ（インデックス付き） |
| `created_at` | TIMESTAMP | 作成日時 |
| `updated_at` | TIMESTAMP | 更新日時 |

### インデックス

- `PRIMARY KEY (id)`
- `idx_channels_is_playing (is_playing)`
- `idx_channels_name (name)`

---

## channel_logs

10分間隔で記録されるチャンネルのスナップショット。

| カラム | 型 | 説明 |
|---|---|---|
| `id` | INT UNSIGNED PK AI | 主キー |
| `log_time` | DATETIME | 記録日時（10分単位に切り捨て） |
| `channel_id` | INT UNSIGNED | channels.id への参照 |
| `cid` | VARCHAR(32) | チャンネルID |
| `name` | VARCHAR(255) | チャンネル名 |
| `bitrate` ~ `track_contact` | | channels と同一スキーマ |
| `hidden_listeners` | TINYINT(1) | リスナー非表示フラグ |
| `created_at` | TIMESTAMP | 作成日時 |
| `updated_at` | TIMESTAMP | 更新日時 |

### インデックス

- `PRIMARY KEY (id)`
- `uix_channel_logs_log_time_name (log_time, name)` — ユニーク（同一時刻・同一チャンネル名の重複防止）
- `idx_channel_logs_channel_id (channel_id)`

---

## channel_daily_summaries

日次バッチで生成される日別集計。

| カラム | 型 | 説明 |
|---|---|---|
| `id` | INT UNSIGNED PK AI | 主キー |
| `log_date` | DATE | 集計対象日 |
| `name` | VARCHAR(255) | チャンネル名 |
| `num_logs` | INT | ログ数（1日の記録回数） |
| `max_listeners` | INT | 最大リスナー数 |
| `average_listeners` | DOUBLE | 平均リスナー数 |
| `created_at` | TIMESTAMP | 作成日時 |
| `updated_at` | TIMESTAMP | 更新日時 |

### インデックス

- `PRIMARY KEY (id)`
- `uix_channel_daily_logs_log_date_name (log_date, name)` — ユニーク

---

## information

index.txt 末尾に追加されるお知らせ情報。

| カラム | 型 | 説明 |
|---|---|---|
| `id` | INT UNSIGNED PK AI | 主キー |
| `name` | VARCHAR(255) | タイトル |
| `description` | VARCHAR(255) | 概要 |
| `priority` | INT (default: -10) | 表示優先度 |
| `created_at` | TIMESTAMP | 作成日時 |
| `updated_at` | TIMESTAMP | 更新日時 |
| `deleted_at` | TIMESTAMP NULL | 論理削除日時（GORM soft delete） |

---

## ER図（概念）

```
channels 1──* channel_logs
   │
   │  (nameで関連)
   ▼
channel_daily_summaries

information (独立)
```

`channel_logs.channel_id` は `channels.id` を参照するが、外部キー制約は設定されていない（アプリケーション側で管理）。
