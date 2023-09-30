package loginpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func LoginPageRegister(r *gin.RouterGroup) {
	r.GET("/login", renderLoginPage)
	r.GET("/create-account", renderCreateAccountForm)
}

func renderLoginPage(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLloginpage", nil)
}

func renderCreateAccountForm(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "createaccountformHTMLloginpage", nil)
}
