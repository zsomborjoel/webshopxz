package cartpage

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/common/ctemplate"
	"github.com/zsomborjoel/workoutxz/internal/model/cart"
)

func CartPageRegister(r *gin.RouterGroup) {
	r.GET("", renderCartPage)
}

func renderCartPage(c *gin.Context) {
	csrfToken := csrf.GetToken(c)
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
		"csrfToken":  csrfToken,
	}

	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLcartpage", dataMap)
}
