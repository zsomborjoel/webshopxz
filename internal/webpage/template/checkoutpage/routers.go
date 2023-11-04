package checkoutpage

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common/consts"
	"github.com/zsomborjoel/workoutxz/internal/common/templaterenderer"
	"github.com/zsomborjoel/workoutxz/internal/model/address"
	"github.com/zsomborjoel/workoutxz/internal/model/cart"
)

func CheckoutPageRegister(r *gin.RouterGroup) {
	r.GET("", renderCheckoutPage)
}

func renderCheckoutPage(c *gin.Context) {
	s := session.GetRoot(c)
	sct := s.Get(consts.Cart)
	if sct == nil {
		log.Error().Msg("Checkout page cart was empty")
		return
	}

	cart := sct.(cart.Cart)
	subtotal := cart.CalculateSubtotal()
	shipping := 10 // TODO store it in db

	userId := s.Get(consts.UserId).(string)

	addr, err := address.FindOneByUserId(userId)
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}

	session.SetCsrfTokenCookie(c)

	dataMap := map[string]interface{}{
		"Cart":     cart,
		"Subtotal": subtotal,
		"Shipping": shipping,
		"Total":    subtotal + shipping,
		"Address":  addr,
		"Username": userId,
	}

	templaterenderer.Render(c.Writer, "indexHTMLcheckoutpage", dataMap)
}
