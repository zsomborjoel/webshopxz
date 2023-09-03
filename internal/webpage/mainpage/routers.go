package mainpage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/model/category"
	"github.com/zsomborjoel/workoutxz/internal/model/product"
	"github.com/zsomborjoel/workoutxz/internal/webpage"
)

type Film struct {
	Title    string
	Director string
}

func MainPageRegister(r *gin.RouterGroup) {
	r.GET("", RenderMainPage)
}

func RenderMainPage(c *gin.Context) {
	products, err := product.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	categories, err := category.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	templates, err := webpage.GetTemplates("/mainpage")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataMap := map[string]interface{}{
		"Products":   products,
		"Categories": categories,
	}

	templates.ExecuteTemplate(c.Writer, "indexHTML", dataMap)
}
