package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	log    *log.Logger
	server *http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.GetHTML)
	mux.HandleFunc("POST /upload", handlers.PostConvertedString)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		log:    logger,
		server: serv,
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
