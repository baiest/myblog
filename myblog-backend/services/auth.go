package services

import (
	"net/http"

	"github.com/baiest/myblog/myblog-backend/handlers/bye"
	"github.com/baiest/myblog/myblog-backend/handlers/hello"
	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/gorilla/mux"
)

var helloHandler = hello.NewByeHandler()
var byeHandler = bye.NewByeHandler()

var routes = []models.Route{
	{
		Path:    "",
		Handler: helloHandler.Create,
		Method:  http.MethodGet,
	},
	{
		Path:    "/bye",
		Handler: byeHandler.Create,
		Method:  http.MethodGet,
	},
}

func AuthRouter(path string, r *mux.Router) {
	router := r.PathPrefix(path).Subrouter()

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
