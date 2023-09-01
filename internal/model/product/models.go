package product

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

type Product struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	SKU         string `db:"sku"`
	Price       int    `db:"price"`
	Active      bool   `db:"active"`
}

func FindAll() ([]Product, error) {
	log.Debug().Msg("products.FindAll called")

	db := common.GetDB()
	var p []Product
	err := db.Select(&p,
		`
		SELECT 
			p.id,
			p.name,
			p.description,
			p.sku, 
			p.price, 
			p.active 
		FROM products p
		`)

	if err != nil {
		return p, fmt.Errorf("An error occured in products.FindAll.Select: %w", err)
	}

	return p, nil
}
