package dto

type CreateEmployeeRequest struct {
	Name     string `json:"name" validate:"required"`
	Position string `json:"position" validate:"required"`
	Salary   int    `json:"salary" validate:"required,min=0"`
}
