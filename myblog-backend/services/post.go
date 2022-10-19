package services

import (
	"net/http"

	"github.com/baiest/myblog/myblog-backend/handlers/post"
	"github.com/baiest/myblog/myblog-backend/models"
)

var postHandler = post.NewPostHandler()

var PostRoutes = []models.Route{
	{
		Path:        "",
		Handler:     postHandler.GetAll,
		Method:      http.MethodGet,
		IsProtected: true,
	},
	{
		Path:        "",
		Handler:     postHandler.Create,
		Method:      http.MethodPost,
		IsProtected: true,
	},
	{
		Path:        "/{id}",
		Handler:     postHandler.GetById,
		Method:      http.MethodGet,
		IsProtected: true,
	},
	{
		Path:        "/{id}",
		Handler:     postHandler.Delete,
		Method:      http.MethodDelete,
		IsProtected: true,
	},
}
