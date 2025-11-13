package app

import (
	"github.com/labstack/echo/v4"
	"github.com/masraga/iykra-test/config"
	"github.com/masraga/iykra-test/internal/handler"
)

type Router struct {
	E  *echo.Echo
	Db config.DatabaseInterface
}

func NewRouter(e *echo.Echo, db config.DatabaseInterface) *Router {
	return &Router{E: e, Db: db}
}

func (r *Router) Register() {
	employeeHandler := handler.NewEmployeeHandler(r.E, r.Db)

	api := r.E.Group("/api")
	api.POST("/employee", employeeHandler.CreateEmployee)
}
