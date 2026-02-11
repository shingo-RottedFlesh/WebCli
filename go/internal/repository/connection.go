package repository

import (
	// ドライバを登録するために「_」（ブランクインポート）でインポート
	"database/sql"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// コンストラクタ
func NewConnectDB() (*sql.DB, error) {
	// DB接続情報
	dbURL := os.Getenv("DATABASE_URL")
	// 環境変数で取得できない場合は開発環境用
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@db:5432/appdb?sslmode=disable"
		// dbURL := "host=db user=postgres password=postgres dbname=appdb sslmode=disable"
	}

	//
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, err
	}
	// connectDB()終了時に接続を閉じる
	// defer db.Close()

	// 接続プール（コネクションプール）の設定
	db.SetMaxOpenConns(25)                 // 最大オープン接続数
	db.SetMaxIdleConns(25)                 // アイドル接続数
	db.SetConnMaxLifetime(5 * time.Minute) // 接続の最大生存期間

	return db, nil
}

// func (db *sql.DB) CloseDB() {
// 	db.Close()
// }
