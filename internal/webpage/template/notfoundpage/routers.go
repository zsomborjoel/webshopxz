package notfoundpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common/ctemplate"
)

func RenderNotFoundPage(c *gin.Context) {
	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLnotfoundpage", nil)
}
