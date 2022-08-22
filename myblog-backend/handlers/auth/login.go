package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var data models.AuthRequest
	var errorResponse models.ErrorResponse
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errorResponse.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	user, err := a.userRepository.GetByEmail(data.Email)
	if err != nil {
		errorResponse.Error = "Invalid Credentials"
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		errorResponse.Error = "Invalid Credentials"
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	claim := models.TokenClaims{
		IdUser: user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		errorResponse.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	json.NewEncoder(w).Encode(models.LoginResponse{
		Token: tokenString,
	})
}
