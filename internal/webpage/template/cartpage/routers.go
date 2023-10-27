package cartpage

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/common/ctemplate"
	"github.com/zsomborjoel/workoutxz/internal/model/cart"
)

func CartPageRegister(r *gin.RouterGroup) {
	r.GET("", renderCartPage)
}

func renderCartPage(c *gin.Context) {
	session := session.GetRoot(c)
	cart := session.Get(common.Cart).(cart.Cart)

	subtotal := cart.CalculateSubtotal()
	shipping := 10 // TODO store it in db

	dataMap := map[string]interface{}{
		"Cart":       cart,
		"Subtotal":   subtotal,
		"Shipping":   shipping,
		"Total":      subtotal + shipping,
		"IsMainPage": true,
	}

	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLcartpage", dataMap)
}
