package userservice

import "github.com/codico/boilerplate/db"

type UserOutput struct {
  ID int64 `json:"id"`
  Username string `json:"username"`
}

func getUserOutput(user db.User) *UserOutput {
  return &UserOutput{
    ID: user.ID,
    Username: user.Username,
  }
}
