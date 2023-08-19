package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zsomborjoel/workoutxz/internal/user"
)

type UserClaim struct {
	jwt.RegisteredClaims
	user.User
}

func CreateJWTToken(user user.User) (string, error) {
	key := os.Getenv("JWT_KEY")

	exp := &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
		User: user,
	})

	signedString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("Error creating signed string: %v", err)
	}

	return signedString, nil
}
