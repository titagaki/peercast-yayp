-- チャンネル情報
CREATE TABLE IF NOT EXISTS channels (
  `id`               INT UNSIGNED AUTO_INCREMENT,
  `cid`              VARCHAR(32),  -- チャンネルID
  `name`             VARCHAR(255), -- チャンネル名
  `bitrate`          INT,          -- ビットレート (単位はkbps)
  `content_type`     VARCHAR(255), -- コンテナタイプ (WMV,FLV,MKVなど)
  `listeners`        INT,          -- リスナー数
  `relays`           INT,          -- リレー数
  `age`              INT UNSIGNED, -- 配信時間 (秒数)
  `genre`            VARCHAR(255), -- ジャンル
  `description`      VARCHAR(255), -- 概要
  `url`              VARCHAR(255), -- コンタクトURL
  `comment`          VARCHAR(255), -- 配信者コメント
  `track_artist`     VARCHAR(255), -- トラック アーティスト
  `track_title`      VARCHAR(255), -- トラック タイトル
  `track_album`      VARCHAR(255), -- トラック アルバム
  `track_genre`      VARCHAR(255), -- トラック ジャンル
  `track_contact`    VARCHAR(255), -- トラック コンタクトURL
  `hidden_listeners` TINYINT(1),   -- リスナー非表示か (ジャンルに?を入れることにより、リスナー数非表示となる)
  `tracker_ip`       VARCHAR(53),  -- トラッカーIP
  `tracker_direct`   TINYINT(1),   -- 直接接続の許可
  `is_playing`       TINYINT(1),   -- 配信中かどうか
  `created_at`       TIMESTAMP NULL,
  `updated_at`       TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  INDEX idx_channels_is_playing (`is_playing`),
  INDEX idx_channels_name (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- チャンネルログ
CREATE TABLE IF NOT EXISTS channel_logs (
  `id`               INT UNSIGNED AUTO_INCREMENT,
  `log_time`         DATETIME,     -- 日時
  `channel_id`       INT UNSIGNED, -- channels.idの外部キー
  `cid`              VARCHAR(32),  -- チャンネルID
  `name`             VARCHAR(255), -- チャンネル名
  `bitrate`          INT,          -- ビットレート (単位はkbps)
  `content_type`     VARCHAR(255), -- コンテナタイプ (WMV,FLV,MKVなど)
  `listeners`        INT,          -- リスナー数
  `relays`           INT,          -- リレー数
  `age`              INT UNSIGNED, -- 配信時間 (秒数)
  `genre`            VARCHAR(255), -- ジャンル
  `description`      VARCHAR(255), -- 概要
  `url`              VARCHAR(255), -- コンタクトURL
  `comment`          VARCHAR(255), -- 配信者コメント
  `track_artist`     VARCHAR(255), -- トラック アーティスト
  `track_title`      VARCHAR(255), -- トラック タイトル
  `track_album`      VARCHAR(255), -- トラック アルバム
  `track_genre`      VARCHAR(255), -- トラック ジャンル
  `track_contact`    VARCHAR(255), -- トラック コンタクトURL
  `hidden_listeners` TINYINT(1),   -- リスナー非表示か (ジャンルに?を入れることにより、リスナー数非表示となる)
  `created_at`       TIMESTAMP NULL,
  `updated_at`       TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX uix_channel_logs_log_time_name (`log_time`, `name`),
  INDEX idx_channel_logs_channel_id (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- チャンネル日別集計
CREATE TABLE IF NOT EXISTS channel_daily_summaries (
  `id`                INT UNSIGNED AUTO_INCREMENT,
  `log_date`          DATE,         -- 日付
  `name`              VARCHAR(255), -- チャンネル名
  `num_logs`          INT,          -- ログ数
  `max_listeners`     INT,          -- 最大視聴者数
  `average_listeners` DOUBLE,       -- 平均視聴者数
  `created_at`       TIMESTAMP NULL,
  `updated_at`       TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX uix_channel_daily_logs_log_date_name (`log_date`, `name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- お知らせ
CREATE TABLE IF NOT EXISTS information (
  `id`          INT UNSIGNED AUTO_INCREMENT,
  `name`        VARCHAR(255), -- タイトル
  `description` VARCHAR(255), -- 概要
  `created_at`  TIMESTAMP NULL,
  `updated_at`  TIMESTAMP NULL,
  `deleted_at`  TIMESTAMP NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
