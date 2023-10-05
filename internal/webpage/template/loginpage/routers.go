package loginpage

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func LoginPageRegister(r *gin.RouterGroup) {
	r.GET("/login", renderLoginPage)
	r.GET("/create-account", renderCreateAccountForm)
	r.GET("/change-password", renderChangePasswordForm)
	r.GET("/reset-password", renderResetPasswordForm)
}

func renderLoginPage(c *gin.Context) {
	csrfToken := csrf.GetToken(c)

	dataMap := map[string]interface{}{
		"csrfToken": csrfToken,
	}

	common.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLloginpage", dataMap)
}

func renderCreateAccountForm(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "createaccountformHTMLloginpage", nil)
}

func renderChangePasswordForm(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "changepasswordformHTMLloginpage", nil)
}

func renderResetPasswordForm(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "resetpasswordformHTMLloginpage", nil)
}
