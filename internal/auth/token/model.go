package authtoken

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	jwt.RegisteredClaims
	UserId string
}

func CreateJWTToken(userId string) (string, error) {
	key := os.Getenv("JWT_KEY")

	exp := &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
		UserId: userId,
	})

	signedString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("Error creating signed string: %v", err)
	}

	return signedString, nil
}

func Parse(token string) (*jwt.Token, error) {
	key := os.Getenv("JWT_KEY")

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
}
