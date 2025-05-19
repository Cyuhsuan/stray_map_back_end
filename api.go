package main

import (
	"fmt"
	"net/http"

	"github.com/Cyuhsuan/stray_map_back_end/middleware"
)

// RegisterRoutes 註冊所有 API 路由
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/auth", middleware.AuthMiddleware(authHandler))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!ddddd")
}
