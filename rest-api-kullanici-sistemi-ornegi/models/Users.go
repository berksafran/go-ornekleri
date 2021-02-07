package models

// Users ...
type Users struct {
	ID        int
	Username  string
	FirstName string
	Profile   string
	Interests []Interest
}
