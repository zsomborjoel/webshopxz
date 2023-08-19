package refreshtoken

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/user"
)

type RefreshToken struct {
	Token     string `db:"token"`
	CreatedAt int64  `db:"created_at"`
	ExpiredAt int64  `db:"expired_at"`
	UserId    string `db:"user_id"`
}

const ExpirationTime = time.Hour * 24

func CreateOne(user user.User) (string, error) {
	log.Debug().Msg("refreshtoken.CreateOne called")

	uuid, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("An error occured in refreshtoken.CreateOne.NewV4: %w", err)
	}

	now := time.Now()

	token := RefreshToken{
		Token:     uuid.String(),
		CreatedAt: now.Unix(),
		ExpiredAt: now.Add(ExpirationTime).Unix(),
		UserId:    user.Id,
	}

	db := common.GetDB()
	tx := db.MustBegin()

	st := `INSERT INTO refresh_tokens (token, created_at, expired_at, user_id) 
			VALUES (:token, :created_at, :expired_at, :user_id)`
	_, err = tx.NamedExec(st, &token)
	if err != nil {
		return "", fmt.Errorf("An error occured in refreshtoken.CreateOne.NamedExec: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return "", fmt.Errorf("An error occured in refreshtoken.CreateOne.Commit: %w", err)
	}

	return token.Token, nil
}

func IsValid(token string) (RefreshToken, error) {
	log.Debug().Msg("verificationtokens.IsValid called")

	db := common.GetDB()
	var rt RefreshToken
	err := db.Get(&rt, "SELECT * FROM refresh_tokens WHERE token=$1", token)
	if err != nil {
		return rt, fmt.Errorf("An error occured in refreshtoken.IsValid.Get: %w", err)
	}

	if rt.ExpiredAt < time.Now().Unix() {
		return rt, fmt.Errorf("Refresh token is expired")
	}

	return rt, nil
}

func DeleteOne(token string) {
	log.Debug().Msg("refreshtoken.DeleteOne called")

	db := common.GetDB()
	db.MustExec("DELETE FROM refresh_tokens WHERE token=$1", token)
}
