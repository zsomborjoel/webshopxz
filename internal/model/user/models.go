package user

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/common/db"
	"github.com/zsomborjoel/workoutxz/internal/common/str"
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

	db := db.Get()
	var u User
	err := db.Get(&u, "SELECT id, username, email, password, active FROM users WHERE username=$1", username)
	if err != nil {
		return u, fmt.Errorf("An error occured in users.FindUserByUserName.Select: %w", err)
	}

	return u, nil
}

func FindByUserId(id string) (User, error) {
	log.Debug().Msg("users.FindByUserId called")

	db := db.Get()
	var u User
	err := db.Get(&u, "SELECT id, username, email, password, active FROM users WHERE id=$1", id)
	if err != nil {
		return u, fmt.Errorf("An error occured in users.FindByUserId.Get: %w", err)
	}

	return u, nil
}

func ExistByUserName(username string) bool {
	log.Debug().Msg("users.ExistByUserName called")

	db := db.Get()
	var i int
	err := db.Get(&i, "SELECT 1 FROM users WHERE username=$1", username)
	if err != nil {
		log.Warn().Err(err)
	}

	return i == 1
}

func CreateOne(user User) error {
	log.Debug().Msg("users.CreateOne called")

	db := db.Get()
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

	db := db.Get()
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

func ResetPasswordByUserName(username string) error {
	log.Debug().Msg("users.ResetPasswordByUserName called")

	db := db.Get()
	newpass := str.GenerateRandom(8)
	_, err := db.Exec("UPDATE users SET password=$1 WHERE username=$2", newpass, username)
	if err != nil {
		return fmt.Errorf("An error occured in users.ResetPasswordById.Exec: %w", err)
	}

	return nil
}
