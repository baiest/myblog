package services

import (
	"fmt"

	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router) {
	for _, route := range Routes {
		route.Router(route.Path, r)
		fmt.Println(route.Path)
	}
}

var Routes = []models.RouteHandle{
	{
		Path:   "/auth",
		Router: AuthRouter,
	},
}
