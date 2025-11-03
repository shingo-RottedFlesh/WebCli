[1mdiff --git a/go/Dockerfile b/go/Dockerfile[m
[1mindex 0aeb1de..e63e58c 100644[m
[1m--- a/go/Dockerfile[m
[1m+++ b/go/Dockerfile[m
[36m@@ -17,34 +17,31 @@[m
     # GOOS=linux: ビルド環境がWindows/Macでも、Linux用の実行ファイルを生成[m
     # -o /app/server: /app/server という名前で実行ファイルを出力[m
     # RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/server ./[m
[31m-    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/server ./[m
[31m-[m
[31m-    # ★ デバッグ用: ここでビルドステージの結果を一時的に保持する新しいステージを作成 ★[m
[31m-    # FROM builder AS debug-stage[m
[31m-[m
[31m-    # # デバッグ用のCMDを追加し、コンテナを永遠に起動したままにする[m
[31m-    # CMD ["sleep", "infinity"][m
[32m+[m[32m    RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/server ./[m
 [m
 [m
 # ========================[m
 # 実行環境[m
 # ========================[m
[31m-    FROM alpine:latest[m
[32m+[m[32m    # FROM alpine:latest[m
[32m+[m[32m    # FROM debian:bookworm-slim[m
[32m+[m[32m    FROM scratch[m
 [m
[31m-    WORKDIR /app[m
[32m+[m[32m    # ★ 修正点：実行に必要な依存関係を追加 ★[m
[32m+[m[32m    # RUN apk add --no-cache ca-certificates musl-utils[m
 [m
[32m+[m[32m    WORKDIR /app[m
[32m+[m[41m    [m
     # ビルドステージ(builder)から、コンパイル済みのバイナリ(/app/server)のみをコピー[m
[31m-    COPY --from=builder /app/server /app/server[m
[32m+[m[32m    COPY --from=builder /app/server /usr/local/bin/server[m
[32m+[m[32m    # COPY --from=builder /app/server /app/server[m
 [m
     # (オプション) HTTPS通信やタイムゾーン設定が必要な場合[m
     # RUN apk add --no-cache ca-certificates tzdata[m
[31m-[m
[32m+[m[41m    [m
     # コンテナがリッスンするポートを公開[m
     EXPOSE 8080[m
 [m
[31m-    # 実行権限を付与[m
[31m-    # RUN chmod +x /app/server [m
[31m-[m
[31m-[m
     # コンテナ起動時に実行するコマンド[m
[31m-    CMD ["/app/server"][m
\ No newline at end of file[m
[32m+[m[32m    # ENTRYPOINT ["/app/server"][m
[32m+[m[32m    ENTRYPOINT ["/usr/local/bin/server"][m
