package handlers

import (
	"context"
	"fmt"
	"github.com/dkeohane/yagsy/app/utils/token"
	"net/http"
	"strings"
)

type AuthHandler struct {
	JwtSecret string
}

func (ah *AuthHandler) authHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			authToken := authHeader[1]
			claims, err := token.Parse(authToken, []byte(ah.JwtSecret))

			if err == nil {
				ctx := context.WithValue(r.Context(), "props", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(err.Error()))
			}
		}
	})
}
