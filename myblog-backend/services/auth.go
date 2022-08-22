package services

import (
	"net/http"

	"github.com/baiest/myblog/myblog-backend/handlers/auth"
	"github.com/baiest/myblog/myblog-backend/models"
)

var authHandler = auth.NewAuthHandler()

var AuthRoutes = []models.Route{
	{
		Path:    "/singin",
		Handler: authHandler.SingIn,
		Method:  http.MethodPost,
	},
	{
		Path:    "/login",
		Handler: authHandler.Login,
		Method:  http.MethodPost,
	},
}
