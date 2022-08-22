package services

import (
	"net/http"

	"github.com/baiest/myblog/myblog-backend/handlers/user"
	"github.com/baiest/myblog/myblog-backend/models"
)

var userHandler = user.NewUserHandler()

var UserRoutes = []models.Route{
	{
		Path:        "",
		Handler:     userHandler.GetAll,
		Method:      http.MethodGet,
		IsProtected: true,
	},
}
