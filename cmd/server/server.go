package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/masraga/iykra-test/config"
	"github.com/masraga/iykra-test/internal/app"
	"github.com/spf13/viper"
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
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_NAME"),
	)

	// Initialize Database
	db := config.NewDatabase("postgres", dsn)

	app.NewRouter(s.E, db).Register()

	s.E.Logger.Fatal(s.E.Start(fmt.Sprintf(":%d", s.Port)))
	fmt.Printf("Server started at %s:%d\n", s.Host, s.Port)
}
