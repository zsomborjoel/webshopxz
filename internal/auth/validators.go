package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func validatePassword(existing string, recived string) error {
	err := bcrypt.CompareHashAndPassword([]byte(existing), []byte(recived))
    if err != nil {
        return errors.New("Invalid bcrypt password check")
    }
    
    return nil
}
