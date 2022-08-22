package models

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

type RouteHandle struct {
	Path   string
	Router func(path string, r *mux.Router)
}
