package userservice

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const PASSWORD_SALT = 10

type Credentials struct {
	Username          string `json:"username"`
	PlaintextPassword string `json:"password"`
}

// TODO: Implement
func (c Credentials) validate() error {
	fmt.Println("[WARNING]: credentials.validate() is not implemented.")
	return nil
}

func (c Credentials) getHashedPassword() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(c.PlaintextPassword), PASSWORD_SALT)
}
