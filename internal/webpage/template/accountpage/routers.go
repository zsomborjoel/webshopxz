package accountpage

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/auth"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/model/address"
)

func AccountPageRegister(r *gin.RouterGroup) {
	r.GET("/account", renderAccountPage)
	r.GET("/account-address", renderAccountAddressForm)
}

func renderAccountPage(c *gin.Context) {
	csrfToken := csrf.GetToken(c)
	session := sessions.Default(c)
	userId := session.Get(common.UserId).(string)

	addr, err := address.FindByUserId(userId)
	if err != nil {
		log.Error().Err(err)
	}

	fmt.Println(addr)
	dataMap := map[string]interface{}{
		"Address":   addr,
		"LoggedIn":  auth.IsLoggedIn(c),
		"csrfToken": csrfToken,
	}

	common.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLaccountpage", dataMap)
}

func renderAccountAddressForm(c *gin.Context) {
	common.GetTemplate().ExecuteTemplate(c.Writer, "accountaddressformHTMLaccountpage", nil)
}
