package entity

type User struct {
	ID         int    `db:"id"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	IsAdmin    bool   `db:"is_admin"`
	CustomerID int    `db:"customer_id"`
}
