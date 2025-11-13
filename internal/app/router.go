package app

import (
	"github.com/labstack/echo/v4"
	"github.com/masraga/iykra-test/internal/handler"
)

type Router struct {
	e *echo.Echo
}

func NewRouter(e *echo.Echo) *Router {
	return &Router{e: e}
}

func (r *Router) Register() {
	employeeHandler := handler.NewEmployeeHandler(r.e)

	api := r.e.Group("/api")
	api.POST("/employee", employeeHandler.CreateEmployee)
}