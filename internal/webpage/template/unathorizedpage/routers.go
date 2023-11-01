package unathorizedpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common/templaterenderer"
)

func RenderUnauthorizedPage(c *gin.Context) {
	templaterenderer.Render(c.Writer, "indexHTMLunauthorizedpage", nil)
}
