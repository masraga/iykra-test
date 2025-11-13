package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/masraga/iykra-test/internal/dto"
)

type EmployeeHandler struct {
	e *echo.Echo
}

func NewEmployeeHandler(e *echo.Echo) *EmployeeHandler {
	return &EmployeeHandler{e: e}
}

func (h *EmployeeHandler) CreateEmployee(e echo.Context) error {
	var req *dto.CreateEmployeeRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"msg": "user created successfully",
	})
}