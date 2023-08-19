package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func PingRegister(router *gin.RouterGroup) {
	router.GET("", ping)
	router.GET("/db", pingDb)
}

func ping(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func pingDb(c *gin.Context) {
	db := common.GetDB()

	err := db.Ping()
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
