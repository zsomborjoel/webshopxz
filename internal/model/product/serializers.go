package product

import "github.com/gin-gonic/gin"

type CartProduct struct {
	Id          string
	Name        string
	Description string
	Price       int
	ImageName   string
	Amount      int
}

type ProductSerializer struct {
	C *gin.Context
	Product
}

func (s *ProductSerializer) CartProduct() CartProduct {
	return CartProduct{
		Id:          s.Product.Id,
		Name:        s.Product.Name,
		Description: s.Product.Description,
		Price:       s.Product.Price,
		ImageName:   s.Product.ImageName,
		Amount:      1,
	}
}
