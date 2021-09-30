package models


type User struct {
	Name string
	Email  string
	Password string
	Balance float64
}

func (user *User) Register () *User{

	return user
}

