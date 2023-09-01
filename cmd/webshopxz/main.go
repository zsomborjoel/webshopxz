package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/zsomborjoel/workoutxz/internal/auth"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/email"
	"github.com/zsomborjoel/workoutxz/internal/middleware"
	"github.com/zsomborjoel/workoutxz/internal/model/product"
	"github.com/zsomborjoel/workoutxz/internal/model/user"
	"github.com/zsomborjoel/workoutxz/internal/ping"
	"github.com/zsomborjoel/workoutxz/internal/webpage/mainpage"
)

func main() {
	common.LoadEnvVariables()

	level := os.Getenv("LOG_LEVEL")
	zerolog.SetGlobalLevel(common.LogLevel(level))

	common.Init()

	r := gin.Default()
	r.Use(
		middleware.CORS(),
		middleware.ErrorHandler(),
	)

	v1 := r.Group("")
	// technical
	ping.PingRegister(v1.Group("/ping"))
	auth.AuthRegister(v1.Group("/auth"))
	email.EmailRegister(v1.Group("/email"))

	// model
	user.UsersRegister(v1.Group("/users"))
	product.ProductsRegister(v1.Group("/products"))

	// template
	mainpage.MainPageRegister(v1.Group("/mainpage"))

	r.Run()
}
