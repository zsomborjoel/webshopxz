package loginpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common/templaterenderer"
)

func LoginPageRegister(r *gin.RouterGroup) {
	r.GET("/login", renderLoginPage)
	r.GET("/create-account", renderCreateAccountForm)
	r.GET("/change-password", renderChangePasswordForm)
	r.GET("/reset-password", renderResetPasswordForm)
}

func renderLoginPage(c *gin.Context) {
	session.SetCsrfTokenCookie(c)
	templaterenderer.Render(c.Writer, "indexHTMLloginpage", nil)
}

func renderCreateAccountForm(c *gin.Context) {
	templaterenderer.Render(c.Writer, "createaccountformHTMLloginpage", nil)
}

func renderChangePasswordForm(c *gin.Context) {
	templaterenderer.Render(c.Writer, "changepasswordformHTMLloginpage", nil)
}

func renderResetPasswordForm(c *gin.Context) {
	templaterenderer.Render(c.Writer, "resetpasswordformHTMLloginpage", nil)
}
