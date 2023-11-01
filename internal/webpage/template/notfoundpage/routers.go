package notfoundpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common/templaterenderer"
)

func RenderNotFoundPage(c *gin.Context) {
	templaterenderer.Render(c.Writer, "indexHTMLnotfoundpage", nil)
}
