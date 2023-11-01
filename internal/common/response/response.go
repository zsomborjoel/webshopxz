package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zsomborjoel/workoutxz/internal/common/templaterenderer"
	"github.com/zsomborjoel/workoutxz/internal/webpage/template/unathorizedpage"
)

func AbortWithUnauthorizedHtml(c *gin.Context) {
	unathorizedpage.RenderUnauthorizedPage(c)
	c.AbortWithError(http.StatusUnauthorized, errors.New("Unathorized Request"))
}

func AbortWithHtml(c *gin.Context, code int, msg string) {
	dataMap := map[string]string{
		"ErrorMessage": msg,
	}
	templaterenderer.Render(c.Writer, "errorresponseHTMLgeneral", dataMap)
	c.AbortWithError(code, errors.New(msg))
}

func OkWithHtml(c *gin.Context, msg string) {
	dataMap := map[string]string{
		"OkMessage": msg,
	}
	templaterenderer.Render(c.Writer, "okresponseHTMLgeneral", dataMap)
}

func NoItemsHtml(c *gin.Context, msg string) {
	dataMap := map[string]string{
		"NoItemsMessage": msg,
	}
	templaterenderer.Render(c.Writer, "noitemsHTMLgeneral", dataMap)
}
