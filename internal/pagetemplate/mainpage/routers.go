package mainpage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/pagetemplate"
)

type Film struct {
	Title    string
	Director string
}

func MainPageRegister(r *gin.RouterGroup) {
	r.GET("/", RenderMainPage)
}

func RenderMainPage(c *gin.Context) {
	templates, err := pagetemplate.GetTemplates("/mainpage")

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	templates.ExecuteTemplate(c.Writer, "indexHTML", "")
}
