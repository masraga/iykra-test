package usecase

import (
	"github.com/masraga/iykra-test/internal/domain"
	"github.com/masraga/iykra-test/internal/dto"
	"github.com/masraga/iykra-test/internal/repositories"
)

type EmployeeUsecase interface {
	Create(employeeData *dto.CreateEmployeeRequest) error
}

type employeeUsecase struct {
	employeeRepo repositories.EmployeeRepo
}

func NewEmployeeUsecase(employeeRepo repositories.EmployeeRepo) *employeeUsecase {
	return &employeeUsecase{employeeRepo: employeeRepo}
}

func (u *employeeUsecase) Create(employeeData *dto.CreateEmployeeRequest) error {
	var employee domain.Employee
	employee.Name = employeeData.Name
	employee.Position = employeeData.Position
	employee.Salary = employeeData.Salary
	return u.employeeRepo.Create(&employee)
}
