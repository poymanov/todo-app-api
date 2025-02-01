package middleware

import (
	"context"
	"errors"
	"net/http"
	"poymanov/todo/pkg/jwt"
	"poymanov/todo/pkg/response"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeNotAuth(w http.ResponseWriter) {
	err := errors.New(http.StatusText(http.StatusUnauthorized))
	response.JsonError(w, err, http.StatusUnauthorized)
}

func Auth(next http.Handler, jwt *jwt.JWT) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			writeNotAuth(w)
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		isValid, data := jwt.Parse(token)

		if !isValid {
			writeNotAuth(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
