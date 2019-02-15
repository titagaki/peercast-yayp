
-- チャンネル情報
-- https://github.com/kumaryu/peercaststation/wiki/index.txt%E3%81%AE%E4%BB%95%E6%A7%98
CREATE TABLE IF NOT EXISTS channels (
  id                 BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  channel_id         VARCHAR(50)     NOT NULL,               -- チャンネルID
  channel_name       VARCHAR(100)    NOT NULL,               -- チャンネル名
  tip                VARCHAR(53)     NOT NULL,               -- トラッカーIP
  bitrate            INT             NOT NULL DEFAULT 0,     -- ビットレート (単位はkbps)
  container_type     VARCHAR(20)     NOT NULL DEFAULT 'UNKNOWN', -- コンテナタイプ (WMV,FLV,MKVなど)
  listeners          INT             NOT NULL DEFAULT 0,     -- リスナー数
  relays             INT             NOT NULL DEFAULT 0,     -- リレー数
  age                INT             NOT NULL DEFAULT 0,     -- 配信時間 (秒数)
  genre              VARCHAR(50),                            -- ジャンル
  description        VARCHAR(255),                           -- 概要
  url                VARCHAR(255),                           -- コンタクトURL
  comment            VARCHAR(255),                           -- 配信者コメント
  track_artist       VARCHAR(100),                           -- トラック アーティスト
  track_title        VARCHAR(100),                           -- トラック タイトル
  track_album        VARCHAR(100),                           -- トラック アルバム
  track_genre        VARCHAR(100),                           -- トラック ジャンル
  track_contact      VARCHAR(100),                           -- トラック コンタクトURL
  host_direct        ENUM('yes', 'no') NOT NULL DEFAULT 'no', -- 直接接続の許可
  status             VARCHAR(5),                             -- "click"という文字列
  ns                 VARCHAR(20),                            -- 名前空間 (NameSpace:Genreの形で入力すると名前空間として使える)
  display_listeners  ENUM('show', 'hide') NOT NULL DEFAULT 'show', -- リスナー数表示/非表示 (ジャンルに?を入れることにより、リスナー数非表示となる)
  limit_type         ENUM('none', 'port0', 'speed', 'high_speed') NOT NULL DEFAULT 'none', -- 表示制限 (ジャンルに@を入れることにより、帯域制限やポート0の視聴制限を行える)
  permission         ENUM('allowed', 'denied') NOT NULL DEFAULT 'allowed', -- 掲載可否
  is_finished        ENUM('yes', 'no') NOT NULL DEFAULT 'no',              -- 終了しているか
  created            DATETIME NOT NULL,
  updated            DATETIME NOT NULL,
  PRIMARY KEY (id),
  INDEX idx_channel_id (channel_id),
  INDEX idx_channel_name (channel_name),
  INDEX idx_permission_is_finished (permission, is_finished)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- チャンネルログ
CREATE TABLE IF NOT EXISTS channel_logs (
  id                 BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  log_datetime       DATETIME        NOT NULL,               -- 日時
  channel_id         VARCHAR(50)     NOT NULL,               -- チャンネルID
  channel_name       VARCHAR(100)    NOT NULL,               -- チャンネル名
  bitrate            INT             NOT NULL DEFAULT 0,     -- ビットレート (単位はkbps)
  container_type     VARCHAR(20)     NOT NULL DEFAULT 'UNKNOWN', -- コンテナタイプ (WMV,FLV,MKVなど)
  listeners          INT             NOT NULL DEFAULT 0,     -- リスナー数
  relays             INT             NOT NULL DEFAULT 0,     -- リレー数
  age                INT             NOT NULL DEFAULT 0,     -- 配信時間 (秒数)
  genre              VARCHAR(50),                            -- ジャンル
  description        VARCHAR(255),                           -- 概要
  url                VARCHAR(255),                           -- コンタクトURL
  comment            VARCHAR(255),                           -- 配信者コメント
  track_artist       VARCHAR(100),                           -- トラック アーティスト
  track_title        VARCHAR(100),                           -- トラック タイトル
  track_album        VARCHAR(100),                           -- トラック アルバム
  track_genre        VARCHAR(100),                           -- トラック ジャンル
  track_contact      VARCHAR(100),                           -- トラック コンタクトURL
  display_listeners  ENUM('show', 'hide') NOT NULL DEFAULT 'show', -- リスナー数表示/非表示 (ジャンルに?を入れることにより、リスナー数非表示となる)
  PRIMARY KEY (id),
  UNIQUE INDEX idx_log_datetime_channel_name (log_datetime, channel_name),
  INDEX idx_channel_id (channel_id),
  INDEX idx_channel_name (channel_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- チャンネルログ日付
CREATE TABLE IF NOT EXISTS channel_log_dates (
  id                 BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  log_date           DATE            NOT NULL,               -- 日付
  channel_name       VARCHAR(100)    NOT NULL,               -- チャンネル名
  PRIMARY KEY (id),
  UNIQUE INDEX idx_log_date_channel_name (log_date, channel_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
