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
	loginpageTemplateFiles := webpage.GetTemplateFiles("/loginpage")
	if len(loginpageTemplateFiles) == 0 {
		log.Fatal().Stack().Msg("Error loading loginpageTemplateFiles")
	}

	componentTemplateFiles := webpage.GetTemplateFiles("/component")
	if len(componentTemplateFiles) == 0 {
		log.Fatal().Stack().Msg("Error loading loginpageTemplates.componentTemplates")
	}

	mainpageTemplateFiles := webpage.GetTemplateFiles("/mainpage")
	if len(mainpageTemplateFiles) == 0 {
		log.Fatal().Stack().Msg("Error loading loginpageTemplates.mainpageTemplates")
	}

	loginpageTemplates, err = template.New("loginpage").ParseFiles(loginpageTemplateFiles...)
	if err != nil {
		log.Fatal().Stack().Msg("Error loading loginpageTemplates")
	}

	loginpageTemplates.ParseFiles(componentTemplateFiles...)
	loginpageTemplates.ParseFiles(mainpageTemplateFiles...)
}

func LoginPageRegister(r *gin.RouterGroup) {
	r.GET("/login", RenderLoginPage)
}

func RenderLoginPage(c *gin.Context) {
	loginpageTemplates.ExecuteTemplate(c.Writer, "indexHTMLloginpage", nil)
}
