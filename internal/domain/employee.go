package domain

type Employee struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Position string `db:"position"`
	Salary   int    `db:"salary"`
}
