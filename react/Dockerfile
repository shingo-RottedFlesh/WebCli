# ========================
# 🚀 ステージ 1: ビルド環境 (Node.js)
# ========================

FROM node:20-alpine AS builder

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係のインストール (キャッシュを活用するため package.json/lockを先にコピー)
COPY package.json package-lock.json ./
RUN npm install

# ソースコードをすべてコピー
COPY . .

# Reactアプリケーションをビルド (通常、/app/build に静的ファイルが生成されます)
RUN npm run build


# ========================
# 📦 ステージ 2: 実行環境 (Apache)
# ========================
# 軽量なAlpineベースのApacheイメージを使用
FROM httpd:2.4-alpine

# Apacheのデフォルトの公開ディレクトリ（DocumentRoot）は /usr/local/apache2/htdocs/ です

# ビルドステージから生成された静的ファイルを実行環境のApacheの公開ディレクトリにコピー
COPY --from=builder /app/dist /usr/local/apache2/htdocs/

# (オプション) ReactのSPAに対応するためのカスタム設定ファイルをコピーする場合
# 例: COPY ./apache/httpd.conf /usr/local/apache2/conf/httpd.conf

# コンテナがリッスンするポートを公開 (Apacheのデフォルトポート)
EXPOSE 80

# コンテナ起動時にApacheを実行 (httpd:alpineイメージのデフォルトCMDが自動実行されます)
# CMD ["httpd", "-D", "FOREGROUND"] # <-- デフォルトCMDのため、記述不要