package models

type UserBase struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Tg      string `json:"tg"`
}

type User struct {
	ID int `json:"ID"`
	UserBase
}

type UserCreate struct {
	UserBase
	PWD string `json:"PWD"`
}

type UserLogin struct {
	Email string `json:"email"`
	PWD   string `json:"PWD"`
}
