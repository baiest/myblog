package services

import (
	"fmt"

	"github.com/baiest/myblog/myblog-backend/middlewares"
	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router) {
	for _, route := range Routes {
		route.Router(r.PathPrefix(route.Path).Subrouter())
		fmt.Println(route.Path)
	}
}

func AddRouter(routes []models.Route) func(r *mux.Router) {
	return func(r *mux.Router) {
		for _, route := range routes {
			handler := route.Handler
			if route.IsProtected {
				handler = middlewares.CheckToken(handler)
			}
			r.HandleFunc(
				route.Path,
				handler,
			).Methods(route.Method)
		}
	}
}

var Routes = []models.RouteHandle{
	{
		Path:   "/auth",
		Router: AddRouter(AuthRoutes),
	},
	{
		Path:   "/users",
		Router: AddRouter(UserRoutes),
	},
}

func init() {

}
