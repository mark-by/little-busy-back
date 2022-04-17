package entity

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
