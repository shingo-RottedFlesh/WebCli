package main

import (
	repository "app/internal/repository"
	"fmt"
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

		// レスポンス用の構造体
		type LoginResponse struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}

		var u user
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("username：%v\n", u.UserName)
		fmt.Printf("password：%v\n", u.Password)

		userRepo := repository.NewUserRepository(db)
		isValid, err := userRepo.ValidatePassword(u.UserName, u.Password)

		var resp LoginResponse
		if err != nil {
			// エラーが発生した場合（DB接続エラーなど）
			log.Printf("Login error: %v", err)
			resp = LoginResponse{
				Success: false,
				Message: "システムエラーが発生しました",
			}
			w.WriteHeader(http.StatusInternalServerError)
		} else if !isValid {
			// 認証失敗
			fmt.Printf("認証失敗\n")
			resp = LoginResponse{
				Success: false,
				Message: "ユーザ名またはパスワードが間違っています",
			}
			// 200 OK で返しつつ success: false にするのが一般的だが、
			// ここでは要件に合わせてステータスコードを使い分けるか、
			// シンプルに200で統一してJSONで判断させるか。
			// フロントエンドの変更に合わせて200 OKで返却し、JSONの中身で判定させる実装にします。
			w.WriteHeader(http.StatusOK)
		} else {
			// 認証成功
			fmt.Printf("成功\n")
			resp = LoginResponse{
				Success: true,
				Message: "ログイン成功",
			}
			w.WriteHeader(http.StatusOK)
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Failed to encode response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
