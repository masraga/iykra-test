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

	r.E.POST("/employees", employeeHandler.CreateEmployee)
	r.E.GET("/employees", employeeHandler.GetAllEmployees)
	r.E.GET("/employees/:id", employeeHandler.GetEmployeeByID)
	r.E.PUT("/employees/:id", employeeHandler.UpdateEmployee)
	r.E.DELETE("/employees/:id", employeeHandler.DeleteEmployee)
}
