package models

type User struct {
	ID       int    `json: "id"`
	Name     string `json: "name"`
	Age      int    `json: "Age"`
	Address  string `json: "Address"`
	UserType int    `json: "UserType"`
}

