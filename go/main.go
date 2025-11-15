package main

import (
	"app/dbdriver"
	"fmt"
	"log"
	"net/http"

	"encoding/json"
)

type result struct {
	Text string `json:text`
}

func main() {
	// 起動時に1度だけDB接続を初期化
	db, err := dbdriver.NewConnectDB()
	if err != nil {
		log.Fatalf("FATAL: Failed to initialize database: %v", err)
	}
	// 2. アプリケーション終了時に**一度だけ**Closeを呼ぶ
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result{Text: "testだぜ！！！"})

		userRepo := dbdriver.NewUserRepository(db)
		user, err := userRepo.ValidatePassword("shingo", "shingo")

		// hash, err := bcrypt.GenerateFromPassword([]byte("shingo"), bcrypt.DefaultCost)
		if err != nil && !user {
			fmt.Printf("認証失敗\n")
		} else {
			fmt.Printf("成功\n")
		}

		// insUser := dbdriver.NewUser("shingo", "shingo")
		// err = userRepo.InsertUser(*insUser)

		fmt.Printf("user：%v\n", user)
		fmt.Printf("err：%v\n", err)
	})
	log.Println("Go API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
