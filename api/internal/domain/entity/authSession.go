package entity

type AuthSession struct {
	ID             string
	UserID         int
	ExpirationDate int64
}
