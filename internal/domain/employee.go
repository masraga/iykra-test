package domain

type Employee struct {
	ID int `db:"id"`
	name string `db:"name"`
	position string `db:"position"`
	salary int `db:"salary"`
}