package product

import "github.com/gin-gonic/gin"

type ProductResponse struct {
	Id          string
	Name        string
	Description string
	SKU         string
	Price       int
	Active      bool
}

type ProductSerializer struct {
	C *gin.Context
	Product
}

func (s *ProductSerializer) Response() ProductResponse {
	return ProductResponse{
		Id:          s.Product.Id,
		Name:        s.Product.Name,
		Description: s.Product.Description,
		SKU:         s.Product.SKU,
		Price:       s.Product.Price,
		Active:      s.Product.Active,
	}
}
