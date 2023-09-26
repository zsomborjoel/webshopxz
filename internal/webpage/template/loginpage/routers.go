package loginpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func LoginPageRegister(r *gin.RouterGroup) {
	r.GET("/login", renderLoginPage)
	r.GET("/test", test)
}

func renderLoginPage(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLloginpage", nil)
}

func test(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "test", nil)
}
