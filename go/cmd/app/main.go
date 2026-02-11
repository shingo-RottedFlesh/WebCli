package main

import (
	"app/internal/model"
	repository "app/internal/repository"
	"log"
	"net/http"

	"encoding/json"
)

type result struct {
	Text string `json:"text"`
}

type user struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type file struct {
	FileId   string `json:"file_id"`
	FileName string `json:"file_name"`
}

func main() {
	// 起動時に1度だけDB接続を初期化
	db, err := repository.NewConnectDB()
	if err != nil {
		log.Fatalf("FATAL: Failed to initialize database: %v", err)
	}
	// 2. アプリケーション終了時に**一度だけ**Closeを呼ぶ
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(result{Text: "testだぜ！！！"})
	})

	mux.HandleFunc("/clip/list", func(w http.ResponseWriter, r *http.Request) {

		files := []*file{
			&file{FileId: "1", FileName: "test1"},
			&file{FileId: "2", FileName: "test2"},
			&file{FileId: "3", FileName: "test3"},
		}

		json.NewEncoder(w).Encode(files)
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		model.Login(w, r, db)
	})

	log.Println("Go API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", commonMiddleware(mux)))

}

// ルーティング事前共通処理
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")

		// CORSのPreflightリクエストを処理
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
