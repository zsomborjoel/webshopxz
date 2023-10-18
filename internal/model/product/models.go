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
	ImageName   string `db:"image_name"`
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
			p.image_name,
			p.active 
		FROM products p
		`)

	if err != nil {
		return p, fmt.Errorf("An error occured in products.FindAll.Select: %w", err)
	}

	return p, nil
}

func FindAllByCategory(cn string) ([]Product, error) {
	log.Debug().Msg("products.FindAllByCategory called")

	db := common.GetDB()
	var p []Product

	if cn == common.All {
		return FindAll()
	}

	err := db.Select(&p,
		`
		SELECT 
			p.id,
			p.name,
			p.description,
			p.sku, 
			p.price, 
			p.image_name,
			p.active 
		FROM products p
		JOIN product_categories pc
		ON p.product_category_id = pc.id
		WHERE pc.name = $1
		`, cn)

	if err != nil {
		return p, fmt.Errorf("An error occured in products.FindAllByCategory.Select: %w", err)
	}

	return p, nil
}

func SearchAllBy(text string) ([]Product, error) {
	log.Debug().Msg("products.SearchAllBy called")

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
			p.image_name,
			p.active 
		FROM products p
		JOIN product_categories pc
		ON p.product_category_id = pc.id
		WHERE p.name LIKE '%' || $1 || '%'
		`, text)

	if err != nil {
		return p, fmt.Errorf("An error occured in products.SearchAllBy.Select: %w", err)
	}

	return p, nil
}
