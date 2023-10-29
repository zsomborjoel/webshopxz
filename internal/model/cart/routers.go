package cart

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/common/response"
	"github.com/zsomborjoel/workoutxz/internal/model/product"
)

func CartRegister(r *gin.RouterGroup) {
	r.POST("/add/:product-id", Add)
	r.DELETE("/remove/:product-id", Remove)
}

func NumberOfSessionItems(c *gin.Context) int {
	log.Debug().Msg("cartcheck.NumberOfCartItems called")

	session := session.GetRoot(c)
	sct := session.Get(common.Cart)
	if sct == nil {
		return 0
	}

	cart := sct.(Cart)
	return cart.NumberOfItems()
}

func Add(c *gin.Context) {
	log.Debug().Msg("cart.Add called")

	productId := c.Param("product-id")
	p, err := product.FindOneById(productId)
	if err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("failed to fetch product for cart | [%s]", productId))
	}

	session := session.GetRoot(c)
	ct := session.Get(common.Cart)
	cart := initCart(ct)

	s := product.ProductSerializer{C: c, Product: p}
	cart.AddProduct(s.CartProduct())
	session.Set(common.Cart, cart)

	err = session.Save()
	if err != nil {
		log.Error().Err(err).Msg("Failed to save session in cart.Add")
	}

	response.OkWithHtml(c, "Product been added to cart")
}

func Remove(c *gin.Context) {
	log.Debug().Msg("cart.Remove called")

	productId := c.Param("product-id")

	session := session.GetRoot(c)
	ct := session.Get(common.Cart).(Cart)
	ct.RemoveProductById(productId)
	session.Set(common.Cart, ct)

	err := session.Save()
	if err != nil {
		log.Error().Err(err).Msg("Failed to save session in cart.Remove")
	}

	c.Header(common.HTMXRedirect, common.Cart)
	c.Status(http.StatusOK)
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
