package ping

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common/db"
)

func PingRegister(router *gin.RouterGroup) {
	router.GET("", ping)
	router.GET("/db", pingDb)
	router.GET("/version", version)
}

func ping(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func pingDb(c *gin.Context) {
	db := db.Get()

	err := db.Ping()
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func version(c *gin.Context) {
	v := os.Getenv("VERSION")

	c.JSON(http.StatusOK, gin.H{"version": v})
}
