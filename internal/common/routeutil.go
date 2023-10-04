package common

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetLastPartUrlPath(url string) string {
	pts := strings.Split(url, "/")
	return pts[len(pts)-1]
}

func AbortWithHtml(c *gin.Context, code int, msg string) {
	dataMap := map[string]interface{}{
		"Message": msg,
	}
	GetTemplate().ExecuteTemplate(c.Writer, "errorresponseHTMLgeneral", dataMap)
	c.AbortWithError(code, errors.New(msg))
}

func OkWithHtml(c *gin.Context, msg string) {
	dataMap := map[string]interface{}{
		"Message": msg,
	}
	GetTemplate().ExecuteTemplate(c.Writer, "okresponseHTMLgeneral", dataMap)
}
