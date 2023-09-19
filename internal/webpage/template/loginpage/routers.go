package loginpage

import (
	"html/template"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/webpage"
)

var loginpageTemplates *template.Template

func Init() {
	var err error
	loginpageTemplates, err = webpage.GetTemplates("/loginpage")
	if err != nil {
		log.Fatal().Stack().Msg("Error loading loginpageTemplates")
		return
	}

	componentTemplates := webpage.GetTemplateFiles("/component")
	if len(componentTemplates) == 0 {
		log.Fatal().Stack().Msg("Error loading loginpageTemplates.componentTemplates")
	}

	mainpageTemplates := webpage.GetTemplateFiles("/mainpage")
	if len(mainpageTemplates) == 0 {
		log.Fatal().Stack().Msg("Error loading loginpageTemplates.mainpageTemplates")
	}

	loginpageTemplates.ParseFiles(componentTemplates...)
	loginpageTemplates.ParseFiles(mainpageTemplates...)
}

func LoginPageRegister(r *gin.RouterGroup) {
	r.GET("/login", RenderLoginPage)
}

func RenderLoginPage(c *gin.Context) {
	loginpageTemplates.ExecuteTemplate(c.Writer, "indexHTMLloginpage", nil)
}
