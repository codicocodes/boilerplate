package userservice

import "golang.org/x/crypto/bcrypt"

const PASSWORD_SALT = 10

type UserInput struct {
  Username string `json:"username"`
  PlaintextPassword string `json:"password"`
}

// TODO: Implement
func (i UserInput) validate() error {
  return nil
}

func (i UserInput) getHashedPassword() ([]byte, error) {
  return bcrypt.GenerateFromPassword([]byte(i.PlaintextPassword), PASSWORD_SALT)
}
