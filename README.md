# PeerCast YAYP

PeerCast Yet Another Yellow Pages (仮)

## 概要

Goで書かれたシンプルなPeerCast YPサーバです。

[YP4G](https://mosax.sakura.ne.jp/yp4g/fswiki.cgi?page=YP4G)と比べて、以下の機能がありません。

- ポートチェック
- 帯域チェック
- 名前空間
- チャット機能

## ビルド方法

```sh
dep ensure -v
go build -v -a -o bin/yayp
```

Dockerを使う場合

```sh
docker build --squash -t peercast-yayp .
```

動作には MySQL と PeerCast が必要です。


### PeerCastの設定

Root Modeで起動して、いい感じに設定してください。（よく知りません）

- Server -> Mode : `Root`

- Root Mode -> Host Update (sec) : `120`(default)

  YP4Gのreadme.txtには60に設定とあった。
  実装を調べると、この間隔でPCP_ROOTパケットを投げて、root情報の更新とトラッカーにPCP_BCSTパケット送信を促すのをやっている。
  なんだけど、そもそもPeerCastStationではPCP_ROOTパケットは破棄するので、今では重要な設定値ではない。

### MySQLの設定

`docker-compose/sql` にスキーマなどがあります。
