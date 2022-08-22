package middlewares

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/golang-jwt/jwt"
)

func CheckToken(next http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var auth string = r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: "Unauthorized",
			})
			return
		}
		tokenString := strings.Split(auth, " ")[1]
		token, err := jwt.ParseWithClaims(tokenString, &models.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Error: "Unauthorized",
			})
			return
		}

		if claim, ok := token.Claims.(*models.TokenClaims); ok && token.Valid {
			r.Header.Add("IdUser", strconv.FormatUint(uint64(claim.IdUser), 10))
			next.ServeHTTP(w, r)
		}
	}
}
