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

func (h *EmployeeHandler) GetAllEmployees(e echo.Context) error {
	employess, err := h.EmployeeUsecase.GetAll()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
	}

	if employess == nil {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"status": http.StatusOK,
			"msg":    "no employees found",
			"data":   []interface{}{},
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   employess,
	})
}

func (h *EmployeeHandler) GetEmployeeByID(e echo.Context) error {
	id := e.Param("id")

	employee, err := h.EmployeeUsecase.GetByID(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   employee,
	})
}

func (h *EmployeeHandler) UpdateEmployee(e echo.Context) error {
	id := e.Param("id")
	var req *dto.UpdateEmployeeRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
	}

	if err := h.EmployeeUsecase.Update(id, req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"msg":    "employee updated successfully",
	})
}

func (h *EmployeeHandler) DeleteEmployee(e echo.Context) error {
	id := e.Param("id")

	if err := h.EmployeeUsecase.Delete(id); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"msg":    "employee deleted successfully",
	})
}
