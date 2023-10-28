package sessionregister

import (
	"encoding/gob"

	"github.com/zsomborjoel/workoutxz/internal/model/cart"
)

func RegisterSlices() {
	gob.Register(cart.Cart{})
}
