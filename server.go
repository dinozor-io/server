package server

import (
	"github.com/dinozor-io/interfaces"
)

type Server struct {
	engine interfaces.Engine
	router interfaces.Router
	port   string
}

func New(engine interfaces.Engine, router interfaces.Router) *Server {
	return &Server{
		engine: engine,
		router: router,
		port:   ":8080",
	}
}

func (s *Server) ChangePort(port string) {
	s.port = port
}

func (s *Server) Router() interfaces.Router {
	return s.router
}

func (s *Server) Port() string {
	return s.port
}

func (s *Server) Serve() {
	s.engine.Init(s)
	s.engine.Run()
}
