package userservice

import (
	"errors"
	"os"
	"time"

	"github.com/codico/boilerplate/db"
	"github.com/golang-jwt/jwt"
)

const JWT_TIME_MINUTES = 60

const MAX_REFRESHES = 100

func getJwtSigningKey() []byte {
	return []byte(os.Getenv("TOKEN_SECRET"))
}

// JWT used to authenticate with and to grant authorization into the system
type JwtToken string

func (t JwtToken) Validate() (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(string(t), claims, func(token *jwt.Token) (interface{}, error) {
		return getJwtSigningKey(), nil
	})
	if err != nil || !tkn.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}

func (t JwtToken) RefreshToken() (*JwtToken, error) {
	claims, err := t.Validate()
	if err != nil {
		return nil, ErrInvalidToken
	}
	if claims.RefreshesRemaining < 1 {
		return nil, ErrInvalidToken
	}
	refreshesRemaining := claims.RefreshesRemaining - 1
	token, err := generateToken(claims.ID, claims.Username, refreshesRemaining)
	return &token, err
}

func (t JwtToken) GetUser() (*UserData, error) {
	return nil, errors.New("not implemented")
}

type Claims struct {
	ID                 int64  `json:"id"`
	Username           string `json:"username"`
	RefreshesRemaining int    `json:"refreshesRemaining"`
	jwt.StandardClaims
}

func generateToken(id int64, username string, refreshesRemaining int) (JwtToken, error) {
	expirationTime := time.Now().Add(JWT_TIME_MINUTES * time.Minute)
	claims := &Claims{
		ID:                 id,
		Username:           username,
		RefreshesRemaining: refreshesRemaining,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJwtSigningKey())
	return JwtToken(tokenString), err
}

func newTokenFromUser(u db.User) (JwtToken, error) {
	token, err := generateToken(u.ID, u.Username, MAX_REFRESHES)
	return token, err
}
