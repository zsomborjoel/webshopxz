package accountpage

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/auth"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func AccountPageRegister(r *gin.RouterGroup) {
	r.GET("/account", renderAccountPage)
}

func renderAccountPage(c *gin.Context) {
	csrfToken := csrf.GetToken(c)

	dataMap := map[string]interface{}{
		"LoggedIn":  auth.IsLoggedIn(c),
		"csrfToken": csrfToken,
	}

	common.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLaccountpage", dataMap)
}
