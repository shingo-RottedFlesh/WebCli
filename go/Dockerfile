# ========================
# ビルド環境
# ========================
    # Go公式のイメージ
    FROM golang:1.25-alpine AS builder

    # 作業Dir作成
    WORKDIR /app

    # 依存関係のキャッシュを活用するため、go.mod と go.sum を先にコピー
    COPY go.mod go.sum ./
    RUN go mod download

    # アプリケーションのソースコードをすべてコピー
    COPY ./firstApp/ .


    # GOOS=linux: ビルド環境がWindows/Macでも、Linux用の実行ファイルを生成
    # -o /app/server: /app/server という名前で実行ファイルを出力
    RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server ./


# ========================
# 実行環境
# ========================
    FROM alpine:latest

    WORKDIR /app

    # ビルドステージ(builder)から、コンパイル済みのバイナリ(/app/server)のみをコピー
    COPY --from=builder /app/server /app/server

    # (オプション) HTTPS通信やタイムゾーン設定が必要な場合
    # RUN apk add --no-cache ca-certificates tzdata

    # コンテナがリッスンするポートを公開
    EXPOSE 8080

    # コンテナ起動時に実行するコマンド
    CMD ["/app/server"]