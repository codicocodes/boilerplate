package userservice

import (
	"context"

	"github.com/codico/boilerplate/db"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
  db db.Queries
  input UserInput
}

func New(db *db.Queries, input UserInput) UserService {
  return UserService{
    db: *db,
    input: input,
  }
}

func (s *UserService) Register() (*UserOutput, error) {
  if err := s.input.validate(); err != nil {
    return nil, newBadUserInputError(err.Error())
	}
  hashedPassword, err := s.input.getHashedPassword()
  if err != nil {
    return nil, newBadUserInputError("password can't be hashed")
	}
  user, err := s.db.CreateUser(
    context.Background(), 
    db.CreateUserParams{
      Username: s.input.Username,
      Password: string(hashedPassword),
    },
  )
  if err != nil {
    return nil, newFailedUserCreationError(err.Error())
	}
  return getUserOutput(user), nil
}

func (s *UserService) Login() (*JwtToken, error) {
  user, err := s.db.GetUserByUsername(context.Background(), s.input.Username)
  if err != nil {
    return nil, ErrLoginFailed
	}
  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(s.input.PlaintextPassword))
  if err != nil {
    return nil, ErrLoginFailed
	}
  token, err := NewTokenFromUser(user)
  if err != nil {
    return nil, ErrLoginFailed
	}
  return &token, nil
}
