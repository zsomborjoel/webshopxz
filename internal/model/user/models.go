package user

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

type User struct {
	Id       string `db:"id"`
	UserName string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Active   bool   `db:"active"`
}

func FindByUserName(username string) (User, error) {
	log.Debug().Msg("users.FindUserByUserName called")

	db := common.GetDB()
	var u User
	err := db.Get(&u, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		return u, fmt.Errorf("An error occured in users.FindUserByUserName.Select: %w", err)
	}

	return u, nil
}

func FindByUserId(id string) (User, error) {
	log.Debug().Msg("users.FindByUserId called")

	db := common.GetDB()
	var u User
	err := db.Get(&u, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return u, fmt.Errorf("An error occured in users.FindByUserId.Get: %w", err)
	}

	return u, nil
}

func ExistByUserName(username string) error {
	log.Debug().Msg("users.ExistByUserName called")

	db := common.GetDB()
	var i int
	err := db.Get(&i, "SELECT 1 FROM users WHERE username=$1", username)
	if err != nil {
		log.Warn().Err(err)
	}

	if i == 1 {
		return errors.New("User already exits")
	}

	return nil
}

func CreateOne(user User) error {
	log.Debug().Msg("users.CreateOne called")

	db := common.GetDB()
	tx := db.MustBegin()

	st := `INSERT INTO users (id, username, email, password) 
			VALUES (:id, :username, :email, :password)`
	_, err := tx.NamedExec(st, &user)
	if err != nil {
		log.Error().Err(err)
		return fmt.Errorf("An error occured in users.CreateOne.NamedExec: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err)
		return fmt.Errorf("An error occured in users.CreateOne.Commit: %w", err)
	}

	return nil
}

func ActivateOne(user User) error {
	log.Debug().Msg("users.ActivateOne called")

	db := common.GetDB()
	tx := db.MustBegin()

	st := `UPDATE users SET active=true WHERE id=:id`
	_, err := tx.NamedExec(st, &user)
	if err != nil {
		return fmt.Errorf("An error occured in users.ActivateOne.NamedExec: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("An error occured in users.ActivateOne.Commit: %w", err)
	}

	return nil
}
