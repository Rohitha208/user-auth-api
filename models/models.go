package models

type User struct {
	Username    string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string // Ensure you hash passwords before saving in production
}
