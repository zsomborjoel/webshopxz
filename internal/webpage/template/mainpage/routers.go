package mainpage

import (
	"net/http"

	"github.com/rs/zerolog/log"
	csrf "github.com/utrack/gin-csrf"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/auth"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/common/ctemplate"
	"github.com/zsomborjoel/workoutxz/internal/model/category"
	"github.com/zsomborjoel/workoutxz/internal/model/product"
	"github.com/zsomborjoel/workoutxz/internal/webpage"
)

var categories []string

func MainPageRegister(r *gin.RouterGroup) {
	r.GET("", renderMainPage)
	r.GET("/search", renderProductsBySearch)
}

func ProductsByCategoryRegister(r *gin.RouterGroup) {
	r.GET("product-categories/:name", renderProductsByCategory)
}

func ProductDetailsByTagNameRegister(r *gin.RouterGroup) {
	r.GET("/product-details/:name", renderProductDetails)
}

func renderProductDetails(c *gin.Context) {
	cats, err := category.FindAllNameWithProducts()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	tag := c.Param("name")
	productByTag, err := product.FindOneByTagName(tag)
	if err != nil {
		log.Warn().Err(err).Msg("productByTag not found in mainpage.renderProductDetails")
	}

	dataMap := map[string]interface{}{
		"Categories": cats,
		"LoggedIn":   auth.IsLoggedIn(c),
		"IsMainPage": true,
		"Product":    productByTag,
	}

	if !webpage.IsHTMXRequest(c) {
		executeMainPage(c, dataMap)
		return
	}

	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "productdetailsHTMLmainpage", dataMap)
}

func renderMainPage(c *gin.Context) {
	ps, err := product.FindAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataMap := map[string]interface{}{
		"Products":  ps,
		"csrfToken": csrf.GetToken(c),
	}

	executeMainPage(c, dataMap)
}

func renderProductsByCategory(c *gin.Context) {
	cat := c.Param("name")
	products, err := product.FindAllByCategory(cat)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataMap := map[string]interface{}{
		"Products": products,
	}

	if !webpage.IsHTMXRequest(c) {
		executeMainPage(c, dataMap)
		return
	}

	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "productHTMLmainpage", dataMap)
}

func renderProductsBySearch(c *gin.Context) {
	t := c.Query("query")

	products, err := product.SearchAllByText(t)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataMap := map[string]interface{}{
		"Products": products,
	}

	if !webpage.IsHTMXRequest(c) {
		executeMainPage(c, dataMap)
		return
	}

	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "productHTMLmainpage", dataMap)
}

func executeMainPage(c *gin.Context, source map[string]interface{}) {
	cats, err := category.FindAllNameWithProducts()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataMap := map[string]interface{}{
		"Categories": cats,
		"LoggedIn":   auth.IsLoggedIn(c),
		"IsMainPage": true,
	}

	common.MergeMaps(source, dataMap)
	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLmainpage", dataMap)
}
