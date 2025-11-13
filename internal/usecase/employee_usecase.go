package usecase

import (
	"github.com/masraga/iykra-test/internal/domain"
	"github.com/masraga/iykra-test/internal/dto"
	"github.com/masraga/iykra-test/internal/repositories"
)

type EmployeeUsecase interface {
	Create(employeeData *dto.CreateEmployeeRequest) error
	GetAll() (*[]domain.Employee, error)
	GetByID(id string) (*domain.Employee, error)
	Update(id string, payload *dto.UpdateEmployeeRequest) error
	Delete(id string) error
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

func (u *employeeUsecase) GetAll() (*[]domain.Employee, error) {
	return u.employeeRepo.GetAll()
}

func (u *employeeUsecase) GetByID(id string) (*domain.Employee, error) {
	return u.employeeRepo.GetByID(id)
}

func (u *employeeUsecase) Update(id string, payload *dto.UpdateEmployeeRequest) error {
	_, err := u.employeeRepo.GetByID(id)
	if err != nil {
		return err
	}
	
	if err := u.employeeRepo.Update(id, payload); err != nil {
		return err
	}
	return nil
}

func (u *employeeUsecase) Delete(id string) error {
	return u.employeeRepo.Delete(id)
}
