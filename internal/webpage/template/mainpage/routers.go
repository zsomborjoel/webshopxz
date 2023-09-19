package mainpage

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/model/category"
	"github.com/zsomborjoel/workoutxz/internal/model/product"
	"github.com/zsomborjoel/workoutxz/internal/webpage"
)

var mainpageTemplates *template.Template
var categories []string

func Init() {
	var err error
	mainpageTemplates, err = webpage.GetTemplates("/mainpage")
	if err != nil {
		log.Fatal().Stack().Msg("Error loading mainpageTemplates")
		return
	}

	componentTemplates := webpage.GetTemplateFiles("/component")
	if len(componentTemplates) == 0 {
		log.Fatal().Stack().Msg("Error loading mainpageTemplates.componentTemplates")
	}

	mainpageTemplates.ParseFiles(componentTemplates...)
}

func MainPageRegister(r *gin.RouterGroup) {
	r.GET("", RenderMainPage)
}

func ProductsByCategoryRegister(r *gin.RouterGroup) {
	cs, err := category.FindAllNameWithProducts()
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

	if !webpage.IsHTMXRequest(c) {
		executeMainPage(c, dataMap)
		return
	}

	mainpageTemplates.ExecuteTemplate(c.Writer, "productHTML", dataMap)
}

func executeMainPage(c *gin.Context, source map[string]interface{}) {
	cats, err := category.FindAllNameWithProducts()
	loggedIn := true
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dataMap := map[string]interface{}{
		"Categories": cats,
		"LoggedIn": loggedIn,
	}

	common.MergeMaps(source, dataMap)
	mainpageTemplates.ExecuteTemplate(c.Writer, "indexHTML", dataMap)
}
