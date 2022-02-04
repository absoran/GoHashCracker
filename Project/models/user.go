package models

type User struct {
	ID       string `json:"id" sql:"id" gorm:"primaryKey"`
	Username string `json:"username" sql:"username"`
	Password string `json:"password" sql:"password"`
	Email    string `json:"email" sql:"email"`
}
