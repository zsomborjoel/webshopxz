package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductsRegister(r *gin.RouterGroup) {
	r.GET("", ProductsRetrieve)
}

func ProductsRetrieve(c *gin.Context) {
	ps, err := FindAll()
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	var prs []ProductResponse
	for _, p := range ps {
		s := ProductSerializer{c, p}
		prs = append(prs, s.Response())
	}

	c.JSON(http.StatusOK, prs)
}
