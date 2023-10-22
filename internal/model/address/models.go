package address

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

type Address struct {
	Id          string `db:"id"`
	Country     string `db:"country"`
	PortalCode  string `db:"postal_code"`
	City        string `db:"city"`
	AddressLine string `db:"address_line"`
	PhoneNumber string `db:"phone_number"`
	CompanyName string `db:"company_name"`
	Details     string `db:"details"`
	UserId      string `db:"user_id"`
	CreatedAt   int64  `db:"created_at"`
	ModifiedAt  int64  `db:"modified_at"`
}

func FindOneByUserId(userId string) (Address, error) {
	log.Debug().Msg("address.FindByUserId called")

	db := common.GetDB()
	var a Address
	err := db.Get(&a,
		`
		SELECT id, country, postal_code, city,
		 address_line, phone_number, company_name, details
		FROM user_addresses 
		WHERE user_id=$1
		`,
		userId)

	if err != nil {
		return a, fmt.Errorf("An error occured in address.FindByUserId.Get: %w", err)
	}

	return a, nil
}

func Update(address Address) error {
	log.Debug().Msg("address.UpdateAddressByUserId called")

	db := common.GetDB()

	address.ModifiedAt = time.Now().Unix()
	_, err := db.NamedExec(
		`UPDATE user_addresses
		SET 
			country=:country,
			postal_code=:postal_code,
			city=:city,
			address_line=:address_line,
			phone_number=:phone_number,
			company_name=:company_name,
			details=:details,
			modified_at=:modified_at
		WHERE user_id=:user_id`,
		address)

	if err != nil {
		return fmt.Errorf("An error occured in address.UpdateAddressByUserId.Exec: %w", err)
	}

	return nil
}

func CreateOne(address Address) error {
	log.Debug().Msg("users.CreateOne called")

	db := common.GetDB()
	tx := db.MustBegin()

	st := `INSERT INTO user_addresses (id, country, postal_code, city, address_line, phone_number, company_name, details, user_id) 
			VALUES (:id, :country, :postal_code, :city, :address_line, :phone_number, :company_name, :details, :user_id)`
	_, err := tx.NamedExec(st, &address)
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

func UpsertOne(address Address) error {
	var err error
	if existByUserId(address.UserId) {
		err = Update(address)
		if err != nil {
			return err
		}

		return nil
	}

	err = CreateOne(address)
	if err != nil {
		return err
	}

	log.Debug().Msg("address.UpsertOne done")
	return nil
}

func existByUserId(id string) bool {
	log.Debug().Msg("address.existByUserId called")

	db := common.GetDB()
	var i int
	err := db.Get(&i, "SELECT 1 FROM user_addresses WHERE user_id=$1", id)
	if err != nil {
		log.Error().Err(err)
	}

	return i == 1
}
