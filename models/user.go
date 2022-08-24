package models

type User struct {
	Id        uint   `jason:"id"`
	FirstName string `jason:"first_name"`
	LastName  string `jason:"last_name"`
	Email     string `jason:"email" gorm:"unique"!`
	Password  []byte `jason:"-"`
}
