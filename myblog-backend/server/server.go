package server

import (
	"context"
	"log"
	"net/http"

	"github.com/baiest/myblog/myblog-backend/db"
	"github.com/baiest/myblog/myblog-backend/middlewares"
	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/baiest/myblog/myblog-backend/services"
	"github.com/gorilla/mux"
)

type Server struct {
	config models.ConfigServer
	router *mux.Router
}

func NewServer(ctx context.Context, config models.ConfigServer) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter().PathPrefix("/api").Subrouter(),
	}
}

func (s *Server) Start() {
	db.NewConnection(s.config.DatabaseConfig)

	s.router.Use(middlewares.SetContentType)
	services.AddRoutes(s.router)

	log.Println("Stating server on port:", s.config.Port)
	if err := http.ListenAndServe(s.config.Port, s.router); err != nil {
		log.Fatal("Fatal error with server:", err.Error())
	}
}
