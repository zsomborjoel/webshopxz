package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"golang.org/x/crypto/bcrypt"
)

func validatePassword(existing string, recived string) error {
	err := bcrypt.CompareHashAndPassword([]byte(existing), []byte(recived))
	if err != nil {
		return errors.New("Invalid bcrypt password check")
	}

	return nil
}

func IsLoggedIn(c *gin.Context) bool {
	session := session.GetRoot(c)
	if at := session.Get(common.AccessToken); at != nil {
		return true
	}

	return false
}
