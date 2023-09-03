package category

import "github.com/gin-gonic/gin"

type CategoryResponse struct {
	Id          string
	Name        string
	Description string
}

type CategorySerializer struct {
	C *gin.Context
	Category
}

func (s *CategorySerializer) Response() CategoryResponse {
	return CategoryResponse{
		Id:          s.Category.Id,
		Name:        s.Category.Name,
		Description: s.Category.Description,
	}
}
