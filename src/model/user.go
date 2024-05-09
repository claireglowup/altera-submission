package model

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	RackId   string `json:"-"`
}
