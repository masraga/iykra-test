package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/masraga/iykra-test/config"
	"github.com/masraga/iykra-test/internal/domain"
	"github.com/masraga/iykra-test/internal/dto"
)

type EmployeeRepo interface {
	Create(employeeData *domain.Employee) error
	GetAll() (*[]domain.Employee, error)
	GetByID(id string) (*domain.Employee, error)
	Update(id string, payload *dto.UpdateEmployeeRequest) error
	Delete(id string) error
}

type employeeRepo struct {
	Db   *sql.DB
	Name string
}

func NewEmployeeRepo(db config.DatabaseInterface) *employeeRepo {
	return &employeeRepo{
		Db:   db.Open(),
		Name: "employees",
	}
}

func (r *employeeRepo) Create(employeeData *domain.Employee) error {
	query := fmt.Sprintf("INSERT INTO %s (name,position,salary) VALUES($1,$2,$3)", r.Name)
	if err := r.Db.QueryRow(query, employeeData.Name, employeeData.Position, employeeData.Salary); err != nil {
		return err.Err()
	}
	return nil
}

func (r *employeeRepo) GetAll() (*[]domain.Employee, error) {
	var employees []domain.Employee
	query := fmt.Sprintf("SELECT id, name, position, salary FROM %s", r.Name)
	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var emp domain.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &employees, nil
}

func (r *employeeRepo) GetByID(id string) (*domain.Employee, error) {
	var emp domain.Employee
	query := fmt.Sprintf("SELECT id, name, position, salary FROM %s WHERE id=$1", r.Name)
	if err := r.Db.QueryRow(query, id).Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary); err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *employeeRepo) Update(id string, payload *dto.UpdateEmployeeRequest) error {
	employee, err := r.GetByID(id)
	if err != nil {
		return err
	}
	if payload.Name == "" {
		payload.Name = employee.Name
	}
	if payload.Position == "" {
		payload.Position = employee.Position
	}
	if payload.Salary == 0 {
		payload.Salary = employee.Salary
	}
	query := fmt.Sprintf("UPDATE %s SET name=$1, position=$2, salary=$3 WHERE id=$4", r.Name)
	if _, err := r.Db.Exec(query, payload.Name, payload.Position, payload.Salary, id); err != nil {
		return err
	}
	return nil
}

func (r *employeeRepo) Delete(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", r.Name)
	if _, err := r.Db.Exec(query, id); err != nil {
		return err
	}
	return nil
}