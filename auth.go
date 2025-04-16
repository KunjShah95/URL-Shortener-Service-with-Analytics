package main

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			sendError(w, http.StatusUnauthorized, "Missing auth token")
			return
		}
		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		
		if err != nil || !token.Valid {
			sendError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		
		next.ServeHTTP(w, r)
	})
}