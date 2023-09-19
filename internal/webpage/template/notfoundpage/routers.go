package notfoundpage

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/webpage"
)

var notfoundpageTemplates *template.Template

func Init() {
	var err error
	notfoundpageTemplates, err = webpage.GetTemplates("/notfoundpage")
	if err != nil {
		log.Fatal().Stack().Msg("Error loading notfoundpageTemplates")
	}

	componentTemplates := webpage.GetTemplateFiles("/component")
	if len(componentTemplates) == 0 {
		log.Fatal().Stack().Msg("Error loading notfoundpageTemplates.componentTemplates")
	}

	notfoundpageTemplates.ParseFiles(componentTemplates...)
}

func RenderNotFoundPage(c *gin.Context) {
	notfoundpageTemplates.ExecuteTemplate(c.Writer, "indexHTMLnotfoundpage", nil)
}
