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
	"github.com/rs/cors"
)

type Server struct {
	config models.ConfigServer
	router *mux.Router
}

func NewServer(ctx context.Context, config models.ConfigServer) *Server {
	r := mux.NewRouter()
	return &Server{
		config: config,
		router: r.PathPrefix("/api").Subrouter(),
	}
}

func (s *Server) Start() {
	db.NewConnection(s.config.DatabaseConfig)
	db.Migrate()

	s.router.Use(middlewares.SetContentType)
	s.router.Use(middlewares.Logger)
	services.AddRoutes(s.router)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	log.Println("Stating server on port:", s.config.Port)
	if err := http.ListenAndServe(s.config.Port, cors.Handler(s.router)); err != nil {
		log.Fatal("Fatal error with server:", err.Error())
	}
}
