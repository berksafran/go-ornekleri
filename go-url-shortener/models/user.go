package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User is model for "/register" handler.
type User struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	CreatedDate time.Time `json:"createdDate"`
}

// Validate is validation of User model.
func (u User) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Username, validation.Required, validation.Length(5, 15)),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

// Validation kısmı kaldı.
