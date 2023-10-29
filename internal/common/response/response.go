package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common/ctemplate"
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
	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "errorresponseHTMLgeneral", dataMap)
	c.AbortWithError(code, errors.New(msg))
}

func OkWithHtml(c *gin.Context, msg string) {
	dataMap := map[string]string{
		"OkMessage": msg,
	}
	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "okresponseHTMLgeneral", dataMap)
}

func NoItemsHtml(c *gin.Context, msg string) {
	dataMap := map[string]string{
		"NoItemsMessage": msg,
	}
	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "noitemsHTMLgeneral", dataMap)
}
