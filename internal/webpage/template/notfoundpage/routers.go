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
	notfoundpageTemplatesFiles := webpage.GetTemplateFiles("/notfoundpage")
	if len(notfoundpageTemplatesFiles) == 0 {
		log.Fatal().Stack().Msg("Error loading notfoundpageTemplatesFiles")
	}

	componentTemplateFiles := webpage.GetTemplateFiles("/component")
	if len(componentTemplateFiles) == 0 {
		log.Fatal().Stack().Msg("Error loading notfoundpageTemplates.componentTemplcomponentTemplateFilesates")
	}

	notfoundpageTemplates, err := template.New("notfoundpage").ParseFiles(notfoundpageTemplatesFiles...)
	if err != nil {
		log.Fatal().Stack().Msg("Error loading notfoundpageTemplates")
	}

	notfoundpageTemplates.ParseFiles(componentTemplateFiles...)
}

func RenderNotFoundPage(c *gin.Context) {
	notfoundpageTemplates.ExecuteTemplate(c.Writer, "indexHTMLnotfoundpage", nil)
}
