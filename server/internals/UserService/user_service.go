package userservice

import (
	"context"
	"errors"

	"github.com/codico/boilerplate/db"
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

func (s *UserService) Login() (*UserOutput, error) {
  return nil, errors.New("not implemented")
}


