package cart

import (
	"github.com/zsomborjoel/workoutxz/internal/model/product"
)

type Cart struct {
	Products map[string]product.CartProduct
}

func EmptyCart() Cart {
	return Cart{Products: make(map[string]product.CartProduct)}
}

func (c *Cart) IsEmpty() bool {
	return len(c.Products) == 0
}

func (c *Cart) NumberOfItems() int {
	return len(c.Products)
}

func (c *Cart) AddProduct(p product.CartProduct) {
	current, isPresent := c.Products[p.Id]
	if !isPresent {
		c.Products[p.Id] = p
		return
	}

	current.Amount = current.Amount + 1
	c.Products[p.Id] = current
}

func (c *Cart) IncreaseProductAmount(productId string) {
	current, isPresent := c.Products[productId]
	if !isPresent {
		return
	}

	current.Amount = current.Amount + 1
	c.Products[productId] = current
}

func (c *Cart) DecreaseProductAmount(productId string) {
	current, isPresent := c.Products[productId]
	if !isPresent {
		return
	}

	if current.Amount == 0 {
		return
	}

	current.Amount = current.Amount - 1
	c.Products[productId] = current
}

func (c *Cart) RemoveProductById(id string) {
	delete(c.Products, id)
}

func (c *Cart) CalculateSubtotal() int {
	sum := 0
	for _, p := range c.Products {
		sum += (p.Price * p.Amount)
	}

	return sum
}
