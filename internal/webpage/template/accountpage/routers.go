package accountpage

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/auth"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/common/ctemplate"
	"github.com/zsomborjoel/workoutxz/internal/model/address"
)

func AccountPageRegister(r *gin.RouterGroup) {
	r.GET("/account", renderAccountPage)
	r.GET("/account-address", renderAccountAddressForm)
}

func renderAccountPage(c *gin.Context) {
	session := session.GetRoot(c)
	userId := session.Get(common.UserId).(string)

	addr, err := address.FindOneByUserId(userId)
	if err != nil {
		log.Error().Err(err)
	}

	dataMap := map[string]interface{}{
		"Address":   addr,
		"LoggedIn":  auth.IsLoggedIn(c),
		"csrfToken": csrf.GetToken(c),
	}

	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLaccountpage", dataMap)
}

func renderAccountAddressForm(c *gin.Context) {
	session := session.GetRoot(c)
	userId := session.Get(common.UserId).(string)

	addr, err := address.FindOneByUserId(userId)
	if err != nil {
		log.Error().Err(err)
	}

	dataMap := map[string]interface{}{
		"Address":   addr,
	}

	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "accountaddressformHTMLaccountpage", dataMap)
}
