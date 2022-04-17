package entity

type Customer struct {
	ID   int64   `db:"id"`
	Name string  `db:"name"`
	Tel  *string `db:"tel"`
}
