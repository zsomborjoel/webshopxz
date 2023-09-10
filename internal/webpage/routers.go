package webpage

import (
	"net/http"
	"text/template"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/model/category"
	"github.com/zsomborjoel/workoutxz/internal/model/product"
)

var mainpageTemplates *template.Template
var componentTemplates *template.Template

func Init() {
	var err error
	mainpageTemplates, err = GetTemplates("/mainpage")
	if err != nil {
		log.Fatal().Stack().Msg("Error loading mainpageTemplates")
		return
	}

	componentTemplates, err = GetTemplates("/component")
	if err != nil {
		log.Fatal().Stack().Msg("Error loading componentTemplates")
	}
}

func MainPageRegister(r *gin.RouterGroup) {
	r.GET("", RenderMainPage)
}

func ProductsByCategoryRegister(r *gin.RouterGroup) {
	cs, err := category.FindAllName()
	if err != nil {
		log.Fatal().Stack().Msg("Error loading ProductsByCategoryRegister routes")
		return
	}

	r.GET(common.AllSlug, RenderProductsByCategory)
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

	dataMap := map[string]interface{}{
		"Products": ps,
	}

	executeMainPage(c, dataMap)
}

func RenderProductsByCategory(c *gin.Context) {
	url := c.Request.URL.String()
	cat := common.GetLastPartUrlPath(url)

	products, err := product.FindAllByCategory(cat)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataMap := map[string]interface{}{
		"Products": products,
	}

	if !IsHTMXRequest(c) {
		executeMainPage(c, dataMap)
		return
	}

	componentTemplates.ExecuteTemplate(c.Writer, "productHTML", dataMap)
}

func executeMainPage(c *gin.Context, source map[string]interface{}) {
	cats, err := category.FindAllName()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataMap := map[string]interface{}{
		"Categories": cats,
	}

	common.MergeMaps(source, dataMap)
	mainpageTemplates.ExecuteTemplate(c.Writer, "indexHTML", dataMap)
}
