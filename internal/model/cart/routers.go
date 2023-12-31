package cart

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common/consts"
	"github.com/zsomborjoel/workoutxz/internal/common/response"
	"github.com/zsomborjoel/workoutxz/internal/model/product"
)

func CartRegister(r *gin.RouterGroup) {
	r.POST("/add/:product-id", Add)
}

func Add(c *gin.Context) {
	log.Debug().Msg("cart.Add called")

	productId := c.Param(consts.ProductId)
	p, err := product.FindOneById(productId)
	if err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("failed to fetch product for cart | [%s]", productId))
	}

	session := session.GetRoot(c)
	ct := session.Get(consts.Cart)
	cart := initCart(ct)

	s := product.ProductSerializer{C: c, Product: p}
	cart.AddProduct(s.CartProduct())
	session.Set(consts.Cart, cart)

	err = session.Save()
	if err != nil {
		log.Error().Err(err).Msg("Failed to save session in cart.Add")
	}

	response.OkWithHtml(c, "Product been added to cart")
}

func IncreaseProductAmount(c *gin.Context) {
	log.Debug().Msg("cart.IncreaseProductAmount called")

	productId := c.Param(consts.ProductId)

	session := session.GetRoot(c)
	ct := session.Get(consts.Cart).(Cart)
	ct.IncreaseProductAmount(productId)
	session.Set(consts.Cart, ct)

	err := session.Save()
	if err != nil {
		log.Error().Err(err).Msg("Failed to save session in cart.IncreaseProductAmount")
	}
}

func DecreaseProductAmount(c *gin.Context) {
	log.Debug().Msg("cart.DecreaseProductAmount called")

	productId := c.Param("product-id")

	session := session.GetRoot(c)
	ct := session.Get(consts.Cart).(Cart)
	ct.DecreaseProductAmount(productId)
	session.Set(consts.Cart, ct)

	err := session.Save()
	if err != nil {
		log.Error().Err(err).Msg("Failed to save session in cart.DecreaseProductAmount")
	}
}

func Remove(c *gin.Context) {
	log.Debug().Msg("cart.Remove called")

	productId := c.Param(consts.ProductId)

	session := session.GetRoot(c)
	ct := session.Get(consts.Cart).(Cart)
	ct.RemoveProductById(productId)
	session.Set(consts.Cart, ct)

	err := session.Save()
	if err != nil {
		log.Error().Err(err).Msg("Failed to save session in cart.Remove")
	}
}

func initCart(ct interface{}) Cart {
	var cart Cart
	if ct == nil {
		cart = EmptyCart()
	} else {
		var ok bool
		cart, ok = ct.(Cart)
		if !ok {
			log.Error().Msg("Failed to assert cart type")
		}
	}
	return cart
}
