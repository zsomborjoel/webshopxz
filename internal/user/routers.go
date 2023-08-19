package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersRegister(r *gin.RouterGroup) {
	r.GET("/:username", UserRetrieveByUserName)
}

func UserRetrieveByUserName(c *gin.Context) {
	un := c.Param("username")
	u, err := FindByUserName(un)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	s := UserSerializer{c, u}
	c.JSON(http.StatusOK, s.Response())
}
