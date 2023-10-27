package cart

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/model/product"
)

func CartRegister(r *gin.RouterGroup) {
	r.POST("/add/:product-id", Add)
	r.DELETE("/remove/:product-id", Remove)
}

func Add(c *gin.Context) {
	log.Debug().Msg("cart.Add called")

	productId := c.Param("product-id")
	p, err := product.FindOneById(productId)
	if err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("failed to fetch product for cart | [%s]", productId))
	}

	session := session.GetRoot(c)
	ct := session.Get(common.Cart).(Cart)
	ct.AddProduct(p)
	session.Set(common.Cart, ct)
	session.Save()
}

func Remove(c *gin.Context) {
	log.Debug().Msg("cart.Remove called")

	productId := c.Param("product-id")

	session := session.GetRoot(c)
	ct := session.Get(common.Cart).(Cart)
	ct.RemoveProductById(productId)
	session.Set(common.Cart, ct)
	session.Save()
}
