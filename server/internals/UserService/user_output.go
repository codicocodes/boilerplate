package userservice

import "github.com/codico/boilerplate/db"

// UserData omitts secret data such as Password
type UserData struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func getUserData(user db.User) *UserData {
	return &UserData{
		ID:       user.ID,
		Username: user.Username,
	}
}
