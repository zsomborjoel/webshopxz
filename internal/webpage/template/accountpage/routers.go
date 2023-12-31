package accountpage

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/auth"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common/consts"
	"github.com/zsomborjoel/workoutxz/internal/common/templaterenderer"
	"github.com/zsomborjoel/workoutxz/internal/model/address"
)

func AccountPageRegister(r *gin.RouterGroup) {
	r.GET("/account", renderAccountPage)
	r.GET("/account-address", renderAccountAddressForm)
}

func renderAccountPage(c *gin.Context) {
	s := session.GetRoot(c)
	userId := s.Get(consts.UserId).(string)

	addr, err := address.FindOneByUserId(userId)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	session.SetCsrfTokenCookie(c)

	dataMap := map[string]interface{}{
		"Address":  addr,
		"LoggedIn": auth.IsLoggedIn(c),
	}

	templaterenderer.Render(c.Writer, "indexHTMLaccountpage", dataMap)
}

func renderAccountAddressForm(c *gin.Context) {
	s := session.GetRoot(c)
	userId := s.Get(consts.UserId).(string)

	addr, err := address.FindOneByUserId(userId)
	if err != nil {
		log.Error().Err(err)
	}

	dataMap := map[string]interface{}{
		"Address": addr,
	}

	templaterenderer.Render(c.Writer, "accountaddressformHTMLaccountpage", dataMap)
}
