package auth

import (
	"errors"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
	session := sessions.Default(c)
	at := session.Get(common.AccessToken)
    fmt.Println(at)
    if at != nil {
        return true
    }
 
    return false
}
