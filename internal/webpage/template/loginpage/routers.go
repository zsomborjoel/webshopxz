package loginpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func LoginPageRegister(r *gin.RouterGroup) {
	r.GET("/login", renderLoginPage)
	r.GET("/create-account", renderCreateAccount)
}

func renderLoginPage(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLloginpage", nil)
}

func renderCreateAccount(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "createaccountHTMLloginpage", nil)
}
