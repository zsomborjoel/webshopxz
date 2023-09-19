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

func FindAllNameWithProducts() ([]Category, error) {
	log.Debug().Msg("categories.FindAll called")

	db := common.GetDB()
	var c []Category
	err := db.Select(&c,
		`
		SELECT pc.name
		FROM product_categories pc
		JOIN products p
		ON pc.id = p.product_category_id
		GROUP BY 1
		`)

	if err != nil {
		return c, fmt.Errorf("An error occured in categories.FindAllNameWithProducts.Select: %w", err)
	}

	return c, nil
}
