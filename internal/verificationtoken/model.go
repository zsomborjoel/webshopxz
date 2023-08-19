package verificationtoken

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/user"
)

type VerificationToken struct {
	Token     string `db:"token"`
	CreatedAt int64  `db:"created_at"`
	ExpiredAt int64  `db:"expired_at"`
	UserId    string `db:"user_id"`
}

const ExpirationTime = time.Hour * 24

func CreateOne(user user.User) (string, error) {
	log.Debug().Msg("verificationtoken.CreateOne called")

	uuid, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("An error occured in verificationtoken.CreateOne.NewV4: %w", err)
	}

	now := time.Now()

	token := VerificationToken{
		Token:     uuid.String(),
		CreatedAt: now.Unix(),
		ExpiredAt: now.Add(ExpirationTime).Unix(),
		UserId:    user.Id,
	}

	db := common.GetDB()
	tx := db.MustBegin()

	st := `INSERT INTO verification_tokens (token, created_at, expired_at, user_id) 
			VALUES (:token, :created_at, :expired_at, :user_id)`
	_, err = tx.NamedExec(st, &token)
	if err != nil {
		return "", fmt.Errorf("An error occured in verificationtoken.CreateOne.NamedExec: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return "", fmt.Errorf("An error occured in verificationtoken.CreateOne.Commit: %w", err)
	}

	return token.Token, nil
}

func IsValid(token string) (VerificationToken, error) {
	log.Debug().Msg("verificationtoken.IsValid called")

	db := common.GetDB()
	var vt VerificationToken
	err := db.Get(&vt, "SELECT * FROM verification_tokens WHERE token=$1", token)
	if err != nil {
		return vt, fmt.Errorf("An error occured in verificationtoken.IsValid.Get: %w", err)
	}

	if vt.ExpiredAt < time.Now().Unix() {
		return vt, fmt.Errorf("Verification token is expired")
	}

	return vt, nil
}

func DeleteOne(token string) {
	log.Debug().Msg("verificationtoken.DeleteOne called")

	db := common.GetDB()
	db.MustExec("DELETE FROM verification_tokens WHERE token=$1", token)
}

func UpdateToken(token string) (string, error) {
	log.Debug().Msg("verificationtoken.UpdateOne called")

	uuid, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("An error occured in verificationtoken.UpdateOne.NewV4: %w", err)
	}

	db := common.GetDB()
	r := db.MustExec("UPDATE verification_tokens SET token=$1 WHERE token=$2", uuid, token)
	num, err := r.RowsAffected() 
	if err != nil {
		return "", fmt.Errorf("An error occured in verificationtoken.UpdateOne.RowsAffected: %w", err)
	}

	if num == 0 {
		return "", errors.New("Token was not updated")
	}

	return uuid.String(), nil	
}
