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
