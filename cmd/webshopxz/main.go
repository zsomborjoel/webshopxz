package main

import (
	"fmt"
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
	common.Init()

	level := os.Getenv("LOG_LEVEL")
	fmt.Println(fmt.Sprintf("LogLevel set to %s", level))
	zerolog.SetGlobalLevel(common.LogLevel(level))

	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)

	r := gin.Default()
	r.Use(
		middleware.CORS(),
		middleware.ErrorHandler(),
		middleware.StaticFileHandler(),
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
