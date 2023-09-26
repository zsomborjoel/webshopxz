package webpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func IsHTMXRequest(c *gin.Context) bool {
	htmx := c.Request.Header.Get(common.HTMXRequest)
	return htmx != ""
}
