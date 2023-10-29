package cartpage

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/common/ctemplate"
	"github.com/zsomborjoel/workoutxz/internal/common/response"
	"github.com/zsomborjoel/workoutxz/internal/model/cart"
)

func CartPageRegister(r *gin.RouterGroup) {
	r.GET("", RenderCartPage)
}

func RenderCartPage(c *gin.Context) {
	noProductMsg := "No product added to cart currently"

	session := session.GetRoot(c)
	sct := session.Get(common.Cart)
	if sct == nil {
		response.NoItemsHtml(c, noProductMsg)
		return
	}

	cart := sct.(cart.Cart)
	isEmptyCart := cart.IsEmpty()
	if sct == nil {
		response.NoItemsHtml(c, noProductMsg)
		return
	}

	csrfToken := csrf.GetToken(c)
	subtotal := cart.CalculateSubtotal()
	numberOfCartItems := cart.NumberOfItems()
	shipping := 10 // TODO store it in db

	dataMap := map[string]interface{}{
		"Cart":              cart,
		"IsEmptyCart":       isEmptyCart,
		"Subtotal":          subtotal,
		"Shipping":          shipping,
		"Total":             subtotal + shipping,
		"IsMainPage":        true,
		"NumberOfCartItems": numberOfCartItems,
		"csrfToken":         csrfToken,
	}

	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLcartpage", dataMap)
}
