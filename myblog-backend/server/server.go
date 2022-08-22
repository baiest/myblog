package server

import (
	"log"
	"net/http"

	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/gorilla/mux"
)

type Server struct {
	config models.ConfigServer
	router *mux.Router
}

func NewServer(config models.ConfigServer) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter().PathPrefix("/api").Subrouter(),
	}
}

func (s *Server) Start() {
	log.Println("Stating server on port:", s.config.Port)
	if err := http.ListenAndServe(s.config.Port, s.router); err != nil {
		log.Fatal("Fatal error with server:", err.Error())
	}
}
