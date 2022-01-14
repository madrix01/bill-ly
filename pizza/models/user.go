package models

type UserLogin struct {
	Username string
	Password string
}

type UserRegister struct {
	Username string
	Password string
	Email    string
}

type User struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
