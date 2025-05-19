package middleware

import (
	"net/http"
)

// make a simple middleware
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	check := false
	return func(w http.ResponseWriter, r *http.Request) {
		if check {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	}
}
