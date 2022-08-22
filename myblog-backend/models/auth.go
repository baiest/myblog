package models

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	IdUser IdUser
	jwt.StandardClaims
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
