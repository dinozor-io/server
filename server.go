package gincon

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	gin    *gin.Engine
	Port   string
	Routes []Route
	Init   func()
}

func (s *Server) Serve() {
	s.LoadDefault()
	s.Init()
	s.gin.Run(s.Port)
}

func (s *Server) LoadDefault() {
	if s.Port == "" {
		s.Port = ":8080"
	}

	if s.gin == nil {
		s.gin = gin.New()
	}

	if s.Init == nil {
		s.Init = func() {
			for _, r := range s.Routes {
				cont := r.Init()
				switch r.Method {
				case GET:
					s.gin.GET(r.Path, cont.Callback)
				case POST:
					s.gin.POST(r.Path, cont.Callback)
				case PUT:
					s.gin.PUT(r.Path, cont.Callback)
				case DELETE:
					s.gin.DELETE(r.Path, cont.Callback)
				}
			}
		}
	}
}
