package userservice

import (
	"context"

	"github.com/codico/boilerplate/db"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db    db.Queries
	credentials Credentials
}

func NewAuth(db *db.Queries, credentials Credentials) UserService {
	return UserService{
		db:    *db,
		credentials: credentials,
	}
}

func (s *UserService) Register() (*UserData, error) {
	if err := s.credentials.validate(); err != nil {
		return nil, newBadUserInputError(err.Error())
	}
	hashedPassword, err := s.credentials.getHashedPassword()
	if err != nil {
		return nil, newBadUserInputError("password can't be hashed")
	}
	user, err := s.db.CreateUser(
		context.Background(),
		db.CreateUserParams{
			Username: s.credentials.Username,
			Password: string(hashedPassword),
		},
	)
	if err != nil {
		return nil, newFailedUserCreationError("user already exists")
	}
	return getUserData(user), nil
}

func (s *UserService) Login() (*JwtToken, error) {
	user, err := s.db.GetUserByUsername(context.Background(), s.credentials.Username)
	if err != nil {
		return nil, ErrLoginFailed
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(s.credentials.PlaintextPassword))
	if err != nil {
		return nil, ErrLoginFailed
	}
	token, err := newTokenFromUser(user)
	if err != nil {
		return nil, ErrLoginFailed
	}
	return &token, nil
}
