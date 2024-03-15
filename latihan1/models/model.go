package models

type User struct {
	ID       int    `json: "id"`
	Name     string `json: "name"`
	Age      int    `json: "Age"`
	Address  string `json: "Address"`
	UserType int    `json: "UserType"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json: "data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json: "data"`
}


