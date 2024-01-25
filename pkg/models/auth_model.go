package models

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type RegisterDto struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=32"`
	FirstName string `json:"firstName" validate:"required,min=2,max=32"`
	LastName  string `json:"lastName" validate:"required,min=2,max=32"`
}
