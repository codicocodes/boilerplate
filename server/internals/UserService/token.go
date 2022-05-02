package userservice

import (
	"os"
	"time"

	"github.com/codico/boilerplate/db"
	"github.com/golang-jwt/jwt"
)

const JWT_TIME_MINUTES = 10

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
	if err != nil {
     return nil, ErrLoginFailed
	}
	if !tkn.Valid {
     return nil, ErrLoginFailed
	}
  return claims, nil
}

type Claims struct {
  ID       int64       `json:"id"`
	Username string      `json:"username"`
	jwt.StandardClaims
}

func NewTokenFromUser(u db.User) (JwtToken, error)  {
   expirationTime := time.Now().Add(JWT_TIME_MINUTES * time.Minute)
   claims := &Claims{
      ID: u.ID,
      Username: u.Username,
      StandardClaims: jwt.StandardClaims{
         IssuedAt: time.Now().Unix(),
         ExpiresAt: expirationTime.Unix(),
      },
   }
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   tokenString, err := token.SignedString(getJwtSigningKey()) 
   return JwtToken(tokenString), err
}
