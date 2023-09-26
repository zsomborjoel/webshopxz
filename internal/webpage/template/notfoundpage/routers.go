package notfoundpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func RenderNotFoundPage(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLnotfoundpage", nil)
}
