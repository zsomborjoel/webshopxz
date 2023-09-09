package webpage

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/model/category"
	"github.com/zsomborjoel/workoutxz/internal/model/product"
)


type Film struct {
	Title    string
	Director string
}

func MainPageRegister(r *gin.RouterGroup) {
	r.GET("", RenderMainPage)
}

func ProductsByCategoryRegister(r *gin.RouterGroup) {
	cs, err := category.FindAllName()
	if err != nil {
		log.Fatal().Stack().Msg("Error loading FilteredProductsByCategoryRegister routes")
	}

	for _, c := range cs {
		r.GET(c.Name, RenderProductsByCategory)
	}
}

func RenderMainPage(c *gin.Context) {
	ps, err := product.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	cats, err := category.FindAllName()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	templates, err := GetTemplates("/mainpage")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataMap := map[string]interface{}{
		"Products":   ps,
		"Categories": cats,
	}

	templates.ExecuteTemplate(c.Writer, "indexHTML", dataMap)
}

func RenderProductsByCategory(c *gin.Context) {
	url := c.Request.URL.String()
	cat := common.GetLastPartUrlPath(url)

	products, err := product.FindAllByCategory(cat)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	templates, err := GetTemplates("/component")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataMap := map[string]interface{}{
		"Products": products,
	}

	templates.ExecuteTemplate(c.Writer, "productHTML", dataMap)
}
