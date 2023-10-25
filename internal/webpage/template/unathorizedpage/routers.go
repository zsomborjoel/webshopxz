package unathorizedpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common/ctemplate"
)

func RenderUnauthorizedPage(c *gin.Context) {
	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLunauthorizedpage", nil)
}
