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

	r.E.POST("/employee", employeeHandler.CreateEmployee)
	r.E.GET("/employee", employeeHandler.GetAllEmployees)
	r.E.GET("/employee/:id", employeeHandler.GetEmployeeByID)
	r.E.PUT("/employee/:id", employeeHandler.UpdateEmployee)
	r.E.DELETE("/employee/:id", employeeHandler.DeleteEmployee)
}
