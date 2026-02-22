# API 仕様

## エンドポイント一覧

| メソッド | パス | 説明 |
|---|---|---|
| GET | `/index.txt` | PeerCast互換のチャンネル一覧（テキスト形式） |
| GET | `/api/channels` | 配信中チャンネル一覧（JSON） |
| GET | `/api/channelLogs` | チャンネルログ（JSON） |
| GET | `/api/channelDailyLogs` | チャンネルログ（`/api/channelLogs` のエイリアス） |
| GET | `/*` | SPAフロントエンド（`public/` から静的配信、HTML5 History Mode対応） |

---

## GET /index.txt

PeerCastクライアントが読み取る標準的な YP index.txt 形式でチャンネル一覧を返す。

### レスポンス

Content-Type: `text/plain`

各チャンネルが1行で、フィールドが `<>` 区切り：

```
チャンネル名<>チャンネルID<>トラッカーIP<>コンタクトURL<>ジャンル<>概要<>リスナー数<>リレー数<>ビットレート<>コンテンツタイプ<>トラックアーティスト<>トラックアルバム<>トラックタイトル<>トラックコンタクト<>URLエンコード済みチャンネル名<>配信時間(H:MM)<>click<>コメント<>直接接続可否(0/1)
```

末尾にお知らせ情報（`information` テーブル）が同形式で追加される。

### リスナー非表示

ジャンル文字列に `?` フラグが含まれるチャンネルは `Listeners=-1`, `Relays=-1` として返される。

---

## GET /api/channels

現在配信中のチャンネル一覧をJSONで返す。

### レスポンス

Content-Type: `application/json`  
CORS: `Access-Control-Allow-Origin: *`

```json
[
  {
    "ID": 1,
    "CID": "0123456789ABCDEF0123456789ABCDEF",
    "Name": "チャンネル名",
    "Bitrate": 1000,
    "ContentType": "FLV",
    "Listeners": 10,
    "Relays": 3,
    "Age": 3600,
    "Genre": "ゲーム",
    "Description": "概要テキスト",
    "Url": "http://example.com",
    "Comment": "配信者コメント",
    "TrackArtist": "",
    "TrackTitle": "",
    "TrackAlbum": "",
    "TrackGenre": "",
    "TrackContact": ""
  }
]
```

### 備考

- キャッシュがあればキャッシュから返却、なければDBからフォールバック
- `HiddenListeners=true` のチャンネルは `Listeners=-1`, `Relays=-1`
- `IsPlaying`, `HiddenListeners`, `TrackerIP`, `TrackerDirect`, `CreatedAt`, `UpdatedAt` はJSONに含まれない

---

## GET /api/channelLogs

指定チャンネル・日付のログをJSONで返す。

### クエリパラメータ

| パラメータ | 必須 | 説明 | 例 |
|---|---|---|---|
| `cn` | Yes | チャンネル名 | `MyChannel` |
| `date` | Yes | 日付（YYYYMMDD形式） | `20240101` |

### レスポンス

Content-Type: `application/json`  
CORS: `Access-Control-Allow-Origin: *`

```json
[
  {
    "ID": 100,
    "LogTime": "2024-01-01T12:00:00+09:00",
    "ChannelID": 1,
    "CID": "0123456789ABCDEF0123456789ABCDEF",
    "Name": "MyChannel",
    "Bitrate": 1000,
    "ContentType": "FLV",
    "Listeners": 10,
    "Relays": 3,
    "Age": 3600,
    "Genre": "ゲーム",
    "Description": "概要",
    "Url": "http://example.com",
    "Comment": "コメント",
    "TrackArtist": "",
    "TrackTitle": "",
    "TrackAlbum": "",
    "TrackGenre": "",
    "TrackContact": ""
  }
]
```

### エラー

| ステータス | 条件 |
|---|---|
| 400 | `date` が8文字でない、または数値変換に失敗 |
| 500 | DB接続失敗など |
