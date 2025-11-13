package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/masraga/iykra-test/config"
	"github.com/masraga/iykra-test/internal/domain"
)

type EmployeeRepo interface {
	Create(employeeData *domain.Employee) error
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
