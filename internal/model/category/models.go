package category

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

type Category struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

func FindAll() ([]Category, error) {
	log.Debug().Msg("categories.FindAll called")

	db := common.GetDB()
	var c []Category
	err := db.Select(&c,
		`
		SELECT 
			pc.id,
			pc.name,
			pc.description
		FROM product_categories pc
		`)

	if err != nil {
		return c, fmt.Errorf("An error occured in categories.FindAll.Select: %w", err)
	}

	return c, nil
}
