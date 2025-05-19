package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	port = ":9000"
)

func init() {
	// 載入 .env 檔案，失敗就直接 panic
	if err := godotenv.Load(); err != nil {
		panic("init error, service stop, reason: " + err.Error())
	}

	envPort := os.Getenv("PORT")
	if envPort != "" {
		// 若 PORT 已經有冒號就不重複加
		if envPort[0] == ':' {
			port = envPort
		} else {
			port = ":" + envPort
		}
	}
}

func main() {
	mux := http.NewServeMux()
	RegisterRoutes(mux)

	fmt.Printf("Listening on port %s\n", port)
	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic("server error: " + err.Error())
	}
}
