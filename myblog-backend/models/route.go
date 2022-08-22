package models

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Path        string
	Handler     http.HandlerFunc
	Method      string
	IsProtected bool
}

type RouteHandle struct {
	Path   string
	Router func(r *mux.Router)
}
