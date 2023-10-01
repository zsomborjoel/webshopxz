package loginpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func LoginPageRegister(r *gin.RouterGroup) {
	r.GET("/login", renderLoginPage)
	r.GET("/create-account", renderCreateAccountForm)
	r.GET("/reset-password", renderResetPasswordForm)
}

func renderLoginPage(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLloginpage", nil)
}

func renderCreateAccountForm(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "createaccountformHTMLloginpage", nil)
}

func renderResetPasswordForm(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "resetpasswordformHTMLloginpage", nil)
}
