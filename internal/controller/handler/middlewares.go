package handler

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nickchervov/go-diplom-project/internal/domain"
)

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		storedPassword := os.Getenv("TODO_PASSWORD")

		if storedPassword == "" {
			next.ServeHTTP(w, r)
			return
		}

		var jwtToken string
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "No cookie with key token", http.StatusUnauthorized)
			return
		}
		jwtToken = cookie.Value

		var claims domain.Claims
		token, err := jwt.ParseWithClaims(jwtToken, &claims, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("incorrect signing method")
			}
			return []byte(storedPassword), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		if claims.PasswordHash != sha256.Sum256([]byte(storedPassword)) {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
