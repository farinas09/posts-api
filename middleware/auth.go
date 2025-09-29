package middleware

import (
	"net/http"
	"strings"

	"github.com/farinas09/rest-ws/models"
	"github.com/farinas09/rest-ws/server"
	"github.com/golang-jwt/jwt/v5"
)

var (
	NO_AUTH_NEEDED = []string{
		"/login",
		"/signup",
	}
)

func shouldCheckToken(url string) bool {
	for _, path := range NO_AUTH_NEEDED {
		if strings.Contains(url, path) {
			return false
		}
	}
	return true
}

func ValidateJWT(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			token := strings.TrimSpace(r.Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(token, &models.AppClaims{}, func(token *jwt.Token) (any, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
