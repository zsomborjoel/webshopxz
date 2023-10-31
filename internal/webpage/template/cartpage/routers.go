package cartpage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/common/ctemplate"
	"github.com/zsomborjoel/workoutxz/internal/common/response"
	"github.com/zsomborjoel/workoutxz/internal/model/cart"
	"github.com/zsomborjoel/workoutxz/internal/webpage"
	"github.com/zsomborjoel/workoutxz/internal/webpage/template/mainpage"
)

func CartPageRegister(r *gin.RouterGroup) {
	r.DELETE("/remove/:product-id", renderRemovedCartItem)
	r.GET("", renderCartBodyPage)
	r.PUT("/increase-product-amount/:product-id", renderCartProductAmountIncrease)
	r.PUT("/decrease-product-amount/:product-id", renderCartProductAmountDecrease)
}

func renderRemovedCartItem(c *gin.Context) {
	cart.Remove(c)
	renderCartBodyPage(c)
}

func renderCartProductAmountIncrease(c *gin.Context) {
	cart.IncreaseProductAmount(c)
	renderCartBodyPage(c)
}

func renderCartProductAmountDecrease(c *gin.Context) {
	cart.DecreaseProductAmount(c)
	renderCartBodyPage(c)
}

func renderCartBodyPage(c *gin.Context) {
	noProductMsg := "No product added to cart currently"

	s := session.GetRoot(c)
	sct := s.Get(common.Cart)
	if sct == nil {
		response.NoItemsHtml(c, noProductMsg)
		return
	}

	cart := sct.(cart.Cart)
	isEmptyCart := cart.IsEmpty()
	if isEmptyCart {
		response.NoItemsHtml(c, noProductMsg)
		return
	}

	subtotal := cart.CalculateSubtotal()
	shipping := 10 // TODO store it in db

	session.SetCsrfTokenCookie(c)

	dataMap := map[string]interface{}{
		"Cart":        cart,
		"IsEmptyCart": isEmptyCart,
		"Subtotal":    subtotal,
		"Shipping":    shipping,
		"Total":       subtotal + shipping,
		"IsMainPage":  true,
	}

	if !webpage.IsHTMXRequest(c) {
		executeMainCartPage(c, dataMap)
		return
	}

	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "bodyHTMLcartpage", dataMap)
}

func executeMainCartPage(c *gin.Context, source map[string]interface{}) {
	dataMap, err := mainpage.GetBaseData(c, source)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctemplate.GetTemplate().ExecuteTemplate(c.Writer, "indexHTMLcartpage", dataMap)
}
