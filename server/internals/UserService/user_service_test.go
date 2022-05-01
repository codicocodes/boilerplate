package userservice_test

import (
	"context"
	"testing"

	testingutils "github.com/codico/boilerplate/internals/TestingUtils"
	userservice "github.com/codico/boilerplate/internals/UserService"
	"golang.org/x/crypto/bcrypt"
)

func setupInputs(t *testing.T, DB testingutils.TestDB) (userservice.UserService, userservice.UserInput){
  DB = testingutils.NewTestConnection()
  input := userservice.UserInput{
    Username: "codico",
    PlaintextPassword: "asdfasdf123",
  }
  service := userservice.New(DB.Queries, input)
  return service, input
}

func TestUserServiceRegisterNoErrors(t *testing.T) {
  // Arrange
  DB := testingutils.NewTestConnection()
  service, _ := setupInputs(t, DB)
  defer DB.Cleanup()

  // Act
  _, err := service.Register()

  // Assert
  if err != nil {
    t.Errorf("Register: expected no errors, actual %s", err.Error())
  }
}

func TestUserServiceRegisterCorrectUsername(t *testing.T) {
  // Arrange
  DB := testingutils.NewTestConnection()
  defer DB.Cleanup()
  service, input := setupInputs(t, DB)

  // Act
  user, err := service.Register()

  // Assert
  if user.Username != input.Username {
    t.Errorf("Register: expected username %s, actual %s", input.Username, err.Error())
  }
}

func TestUserServiceRegisterCorrectPasword(t *testing.T) {
  // Arrange
  DB := testingutils.NewTestConnection()
  service, input := setupInputs(t, DB)
  defer DB.Cleanup()

  // Act
  user, _ := service.Register()

  // Assert
  dbUser, err := DB.Queries.GetUser(context.Background(), user.ID)
  if err != nil {
    panic(err)
  }

  if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.PlaintextPassword)); err != nil {
    t.Errorf("Register: not able to compare hash and password")
  }
}
