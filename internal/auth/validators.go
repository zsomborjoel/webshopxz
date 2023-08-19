package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func validatePassword(recived string, existing string) error {
	err := bcrypt.CompareHashAndPassword([]byte(existing), []byte(recived))
    if err != nil {
        return errors.New("Invalid passowrd")
    }
    
    return nil
}
