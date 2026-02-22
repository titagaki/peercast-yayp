# 1. ビルド用ステージ
FROM golang:1.25-alpine3.21 AS builder

WORKDIR /app

# キャッシュの恩恵を最大化するため、依存関係を先にコピー
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

# キャッシュマウントを維持しつつビルド
# -trimpath: ビルドパスをバイナリから削除（セキュリティ/再現性）
# -ldflags="-w -s": デバッグ情報を削除しサイズ削減
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build \
      -trimpath \
      -ldflags="-w -s" \
      -o /app/yayp ./cmd/yayp

# 2. 実行用ステージ
FROM alpine:3.21

# セキュリティアップデートと最低限のライブラリ
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# builderステージから必要なファイルのみコピー
COPY --from=builder /app/yayp .
COPY --from=builder /app/yayp.toml .
COPY --from=builder /app/public ./public

# 非rootユーザーでの実行（セキュリティ強化が必要な場合はコメント解除）
# RUN addgroup -S appgroup && adduser -S appuser -G appgroup
# USER appuser

EXPOSE 8000
CMD ["./yayp"]
