package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// URL is model of request body.
type URL struct {
	URL          string    `json:"url"`
	Path         string    `json:"path"`
	Description  string    `json:"description"`
	CreatedDate  time.Time `json:"createdDate"`
	VisitedCount int       `json:"visitedCount"`
}

// Validate is validation of request body.
func (u URL) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.URL, validation.Required, is.URL),
		validation.Field(&u.Description, validation.Required, validation.Length(5, 50)),
	)
}
