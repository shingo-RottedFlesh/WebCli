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
    # RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/server ./
    RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/server ./


# ========================
# 実行環境
# ========================
    FROM scratch

    WORKDIR /app

    # ビルドステージ(builder)から、コンパイル済みのバイナリ(/app/server)のみをコピー
    COPY --from=builder /app/server /usr/local/bin/server

    # コンテナがリッスンするポートを公開
    EXPOSE 8080

    # コンテナ起動時に実行するコマンド
    ENTRYPOINT ["/usr/local/bin/server"]
