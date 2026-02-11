package model

import (
	"app/internal/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// レスポンス用の構造体
type loginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("username：%v\n", u.Username)
	fmt.Printf("password：%v\n", u.Password)

	userRepo := repository.NewUserRepository(db)
	isValid, err := userRepo.ValidatePassword(u.Username, u.Password)

	var resp loginResponse
	if err != nil {
		// エラーが発生した場合（DB接続エラーなど）
		log.Printf("Login error: %v", err)
		resp = loginResponse{
			Success: false,
			Message: "システムエラーが発生しました",
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else if !isValid {
		// 認証失敗
		fmt.Printf("認証失敗\n")
		resp = loginResponse{
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
		resp = loginResponse{
			Success: true,
			Message: "ログイン成功",
		}
		w.WriteHeader(http.StatusOK)
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
