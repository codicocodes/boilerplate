package userservice_test

import (
	"context"
	"testing"

	testingutils "github.com/codico/boilerplate/internals/TestingUtils"
	userservice "github.com/codico/boilerplate/internals/UserService"
	"golang.org/x/crypto/bcrypt"
)

func setupRegisterInputs(t *testing.T, DB testingutils.TestDB) (userservice.UserService, userservice.Credentials) {
	input := userservice.Credentials{
		Username:          "codico",
		PlaintextPassword: "Asdfasdf123",
	}
	service := userservice.NewAuth(DB.Queries, input)
	return service, input
}

func TestRegisterNoErrors(t *testing.T) {
	// Arrange
	DB := testingutils.NewTestConnection()
	defer DB.Cleanup()
	service, _ := setupRegisterInputs(t, DB)

	// Act
	_, err := service.Register()

	// Assert
	if err != nil {
		t.Errorf("Register: expected no errors, actual %s", err.Error())
	}
}

func TestRegisterCorrectUsername(t *testing.T) {
	// Arrange
	DB := testingutils.NewTestConnection()
	defer DB.Cleanup()
	service, input := setupRegisterInputs(t, DB)

	// Act
	user, err := service.Register()

	// Assert
	if user.Username != input.Username {
		t.Errorf("Register: expected username %s, actual %s", input.Username, err.Error())
	}
}

func TestRegisterCorrectPasword(t *testing.T) {
	// Arrange
	DB := testingutils.NewTestConnection()
	defer DB.Cleanup()
	service, input := setupRegisterInputs(t, DB)

	// Act
	user, _ := service.Register()

	// Assert
	dbUser, err := DB.Queries.GetUserByID(context.Background(), user.ID)
	if err != nil {
		panic(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.PlaintextPassword)); err != nil {
		t.Errorf("Register: not able to compare hash and password")
	}
}

func TestLoginNoErrors(t *testing.T) {
	// Arrange
	DB := testingutils.NewTestConnection()
	defer DB.Cleanup()
	service, _ := setupRegisterInputs(t, DB)
	_, _ = service.Register()

	// Act
	_, err := service.Login()

	// Assert
	if err != nil {
		t.Errorf("Login: expected no errors, actual %s", err.Error())
	}
}

func TestLoginReceiveToken(t *testing.T) {
	// Arrange
	DB := testingutils.NewTestConnection()
	defer DB.Cleanup()
	service, _ := setupRegisterInputs(t, DB)
	_, _ = service.Register()

	// Act
	token, err := service.Login()

	// Assert
	if err != nil {
		t.Errorf("Login: expected to not receive error, received: %s", err.Error())
	}
	if token == nil {
		t.Error("Login: expected to receive token")
	}
}

func TestLoginWrongPassword(t *testing.T) {
	// Arrange
	DB := testingutils.NewTestConnection()
	defer DB.Cleanup()
	svcRegister, registerInputs := setupRegisterInputs(t, DB)
	_, _ = svcRegister.Register()

	failedInput := userservice.Credentials{
		Username:          registerInputs.Username,
		PlaintextPassword: "not the same password",
	}
	svcLogin := userservice.NewAuth(DB.Queries, failedInput)

	// Act
	token, err := svcLogin.Login()

	// Assert
	if err == nil {
		t.Errorf("Login: expected to receive error, received token %s", string(*token))
	}
}

func TestTokenValidateCorrectClaims(t *testing.T) {
	// Arrange
	DB := testingutils.NewTestConnection()
	defer DB.Cleanup()
	service, _ := setupRegisterInputs(t, DB)
	user, _ := service.Register()
	token, _ := service.Login()

	// Act
	claims, err := token.Validate()

	// Assert
	if err != nil {
		t.Errorf("Token Validate: expected to not receive error, received: %s", err.Error())
	}
	if claims == nil {
		t.Error("Token Validate: expected to receive claims")
	}
	if claims.Username != user.Username {
		t.Errorf("Token Validate: expected to receive username %s, received %s", user.Username, claims.Username)
	}
	if claims.ID != user.ID {
		t.Errorf("Token Validate: expected to receive ID %d, received %d", user.ID, claims.ID)
	}
}

func TestTokenRefreshCorrectRefreshesRemaining(t *testing.T) {
	// Arrange
	expectedRefreshesRemaining := userservice.MAX_REFRESHES - 1
	DB := testingutils.NewTestConnection()
	defer DB.Cleanup()
	service, _ := setupRegisterInputs(t, DB)
	_, _ = service.Register()
	token, _ := service.Login()

	// Act
	token, err := token.RefreshToken()

	// Assert
	if err != nil {
		t.Errorf("Token Refresh: expected to not receive error, received: %s", err.Error())
	}
	claims, _ := token.Validate()

	if claims.RefreshesRemaining != expectedRefreshesRemaining {
		t.Errorf("Token Refresh: expected to receive %d refreshes remaining, received: %d", expectedRefreshesRemaining, claims.RefreshesRemaining)
	}
}
