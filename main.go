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
	_ = godotenv.Load()
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = ":" + envPort
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(port, nil)
}
