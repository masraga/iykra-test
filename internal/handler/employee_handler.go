package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/masraga/iykra-test/config"
	"github.com/masraga/iykra-test/internal/dto"
	"github.com/masraga/iykra-test/internal/repositories"
	"github.com/masraga/iykra-test/internal/usecase"
)

type EmployeeHandler struct {
	e               *echo.Echo
	EmployeeUsecase usecase.EmployeeUsecase
}

func NewEmployeeHandler(e *echo.Echo, db config.DatabaseInterface) *EmployeeHandler {
	employeeRepo := repositories.NewEmployeeRepo(db)
	employeeUsecase := usecase.NewEmployeeUsecase(employeeRepo)
	return &EmployeeHandler{e: e, EmployeeUsecase: employeeUsecase}
}

func (h *EmployeeHandler) CreateEmployee(e echo.Context) error {
	var req *dto.CreateEmployeeRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
	}

	if err := h.EmployeeUsecase.Create(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"msg":    "user created successfully",
	})
}
