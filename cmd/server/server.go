package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/masraga/iykra-test/internal/app"
)

type Server struct {
	Host string
	Port int
	E    *echo.Echo
}

func NewServer(host string, port int) *Server {
	e := echo.New()
	return &Server{
		Host: host,
		Port: port,
		E:    e,
	}
}

func (s *Server) Start() {
	app.NewRouter(s.E).Register()
	s.E.Logger.Fatal(s.E.Start(fmt.Sprintf(":%d", s.Port)))
	fmt.Printf("Server started at %s:%d\n", s.Host, s.Port)
}