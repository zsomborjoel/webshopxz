package webpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common/consts"
)

func IsHTMXRequest(c *gin.Context) bool {
	htmx := c.Request.Header.Get(consts.HTMXRequest)
	return htmx != ""
}
