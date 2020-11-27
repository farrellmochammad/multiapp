package models

type User struct {
	Phone    string `form:"phone", json:"phone"`
	Name     string `form:"name", json:"name"`
	Role     string `form:"role", json:"role"`
	Password string `form:"password", json:"password"`
}

type UserLogin struct {
	Phone    string `form:"phone", json:"phone"`
	Password string `form:"password", json:"password"`
}
