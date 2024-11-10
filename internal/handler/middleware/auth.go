package middleware

import (
	"context"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := &model.Claims{}
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		_, err := jwt.ParseWithClaims(tokenString, claims,
			func(t *jwt.Token) (interface{}, error) {
				return []byte("my_secret_key"), nil
			})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		key := model.ContextKey("login")
		ctx := context.WithValue(r.Context(), key, model.UserLogin{Username: claims.Login})

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}