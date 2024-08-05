package rest

import (
	"net/http"

	"github.com/naysw/permission/internal/core"
)

type Server struct {
	app  *core.App
	mux  *http.ServeMux
	port string
}

type ServerOption func(*Server)

func WithPort(port string) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func defaultServer() *Server {
	return &Server{
		mux:  http.NewServeMux(),
		port: "8080",
		app:  core.NewApp(),
	}
}

func StartServer(app *core.App, opts ...ServerOption) error {
	s := defaultServer()
	for _, opt := range opts {
		opt(s)
	}

	registerRoutes(s.mux, app)

	return http.ListenAndServe(":"+s.port, s.mux)
}
