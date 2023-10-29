package loginpage

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/common/ctemplate"
)

func LoginPageRegister(r *gin.RouterGroup) {
	r.GET("/login", renderLoginPage)
	r.GET("/create-account", renderCreateAccountForm)
	r.GET("/change-password", renderChangePasswordForm)
	r.GET("/reset-password", renderResetPasswordForm)
}

func renderLoginPage(c *gin.Context) {
	dataMap := map[string]interface{}{
		"csrfToken": csrf.GetToken(c),
	}

	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLloginpage", dataMap)
}

func renderCreateAccountForm(c *gin.Context) {
	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "createaccountformHTMLloginpage", nil)
}

func renderChangePasswordForm(c *gin.Context) {
	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "changepasswordformHTMLloginpage", nil)
}

func renderResetPasswordForm(c *gin.Context) {
	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "resetpasswordformHTMLloginpage", nil)
}
