package middlewares

import (
	"net/http"

	"github.com/willchat-ofc/api-willchat-golang/internal/utils"
)

func VerifyAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := w.Header().Get("Authorization")

		if authorization == "" {
			http.Error(w, "Missing or invalid access token", http.StatusUnauthorized)
			return
		}

		_, claims, err := utils.NewCreateAccessTokenUtil().Validate(authorization)

		if err != nil {
			http.Error(w, "Invalid or expired access token", http.StatusUnauthorized)
			return
		}

		r.Header.Set("UserId", claims["sub"].(string))

		next.ServeHTTP(w, r)
	})
}
