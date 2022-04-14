package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/skurtz97/splat/internal/splat"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

type Server struct {
	log     *zap.Logger
	router  *chi.Mux
	service *splat.Service
	server  *http.Server
}

func NewServer(log *zap.Logger, service *splat.Service) *Server {
	log.Info("setting up routes")
	s := &Server{
		log:     log,
		router:  chi.NewRouter(),
		service: service,
	}

	s.router.Use(JSONMiddleware)
	s.router.Use(TimeoutMiddleware)
	s.router.Use(LoggingMiddleware(log))

	s.router.Get("/api/v1/ping", s.Ping)
	s.router.Get("/api/v1/post", s.ListPosts)
	s.router.Post("/api/v1/post", s.CreatePost)
	s.router.Get("/api/v1/post/{id}", s.GetPost)
	s.router.Put("/api/v1/post/{id}", s.UpdatePost)
	s.router.Delete("/api/v1/post/{id}", s.DeletePost)

	log.Info("setting up server")

	s.server = &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
		Addr:         ":8080",
		Handler:      s.router,
	}

	return s
}

func (s *Server) Start() error {
	s.log.Info("starting server")
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			s.log.Error("error starting server", zap.Error(err))
		}
	}()

	// make an interrupt channel and block on it
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// shut down while allowing time for requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.log.Error("error shutting down server", zap.Error(err))
	}

	s.log.Info("server shutdown")
	return nil
}
