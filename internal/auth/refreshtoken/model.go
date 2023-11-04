package refreshtoken

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/common/db"
	"github.com/zsomborjoel/workoutxz/internal/model/user"
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

	db := db.Get()
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

func IsValid(token string) bool {
	log.Debug().Msg("verificationtokens.IsValid called")

	db := db.Get()
	var rt RefreshToken
	err := db.Get(&rt, "SELECT * FROM refresh_tokens WHERE token=$1", token)
	if err != nil {
		log.Error().Err(fmt.Errorf("An error occured in refreshtoken.IsValid.Get: %w", err))
		return false
	}

	if rt.ExpiredAt < time.Now().Unix() {
		log.Debug().Msg("Refresh token expired")
		return false
	}

	return true
}

func DeleteOne(token string) {
	log.Debug().Msg("refreshtoken.DeleteOne called")

	db := db.Get()
	db.Exec("DELETE FROM refresh_tokens WHERE token=$1", token)
}
