package cart

import "github.com/zsomborjoel/workoutxz/internal/model/product"

type Cart struct {
	Products []product.Product
}

func EmptyCart() Cart {
	return Cart{Products: []product.Product{}}
}

func (c *Cart) AddProduct(p product.Product) {
	c.Products = append(c.Products, p)
}

func (c *Cart) RemoveProductById(id string) {
	index := -1
	for i, p := range c.Products {
		if p.Id == id {
			index = i
			break
		}
	}
	if index != -1 {
		c.Products = append(c.Products[:index], c.Products[index+1:]...)
	}
}

func (c *Cart) CalculateSubtotal() int {
	sum := 0
	for _, p := range c.Products {
		sum += p.Price
	}

	return sum
}