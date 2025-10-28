package main

import (
	"log"
	"net/http"

	"encoding/json"
)

type result struct {
	Text string `json:text`
}

func main() {
	// e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	// Reactアプリケーションのオリジンを指定します。
	// 	// "*" を使用すると、任意のオリジンからのリクエストが許可されますが、本番環境では非推奨です。
	// 	AllowOrigins: []string{"*",},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK,
	// 	json.NewEncoder(w).Encode(Message{Text: "testだよ"})
	// 	)
	// })
	// e.Logger.Fatal(e.Start(":8080"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result{Text: "testだぜ"})
	})
	log.Println("Go API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
