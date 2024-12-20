package models

// User represents the structure for the incoming user data
type User struct {
	Name   string `json:"name" validate:"required"`
	PAN    string `json:"pan" validate:"required,pansymbol"`
	Mobile string `json:"mobile" validate:"required,mobile"`
	Email  string `json:"email" validate:"required,email"`
}
