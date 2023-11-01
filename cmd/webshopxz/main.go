package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/zsomborjoel/workoutxz/internal/auth"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/auth/session/sessionregister"
	"github.com/zsomborjoel/workoutxz/internal/common"
	templaterenderer "github.com/zsomborjoel/workoutxz/internal/common/templaterenderer"
	"github.com/zsomborjoel/workoutxz/internal/email"
	"github.com/zsomborjoel/workoutxz/internal/middleware"
	"github.com/zsomborjoel/workoutxz/internal/model/address"
	"github.com/zsomborjoel/workoutxz/internal/model/cart"
	"github.com/zsomborjoel/workoutxz/internal/ping"
	"github.com/zsomborjoel/workoutxz/internal/webpage/template/accountpage"
	"github.com/zsomborjoel/workoutxz/internal/webpage/template/cartpage"
	"github.com/zsomborjoel/workoutxz/internal/webpage/template/loginpage"
	"github.com/zsomborjoel/workoutxz/internal/webpage/template/mainpage"
	"github.com/zsomborjoel/workoutxz/internal/webpage/template/notfoundpage"
)

func main() {
	fmt.Println("Application Init started")
	common.InitEnvVariables()
	common.InitDB()
	templaterenderer.InitTemplates()

	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "error"
	}
	fmt.Println(fmt.Sprintf("LogLevel set to %s", level))
	zerolog.SetGlobalLevel(common.LogLevel(level))

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "release"
	}
	fmt.Println(fmt.Sprintf("GinMode set to %s", ginMode))
	gin.SetMode(ginMode)

	store := session.SetupStore()
	sessionregister.RegisterSlices()

	r := gin.Default()
	r.Use(
		middleware.CORS(),
		sessions.Sessions("mysession", store),
		middleware.CSRFProtectionHandler(),
		middleware.StaticFileHandler(),
		middleware.ErrorHandler(),
		middleware.XSSProtectionHandler(),
	)

	r.NoRoute(notfoundpage.RenderNotFoundPage)

	api := r.Group("/api")
	ping.PingRegister(api.Group("/ping"))
	auth.AuthRegister(api.Group("/auth"))
	email.EmailRegister(api.Group("/email"))

	template := r.Group("")
	mainpage.MainPageRegister(template)
	mainpage.ProductsByCategoryRegister(template)
	mainpage.ProductDetailsByTagNameRegister(template)
	loginpage.LoginPageRegister(template)

	cartgrp := template.Group("/cart")
	cartpage.CartPageRegister(cartgrp)
	cart.CartRegister(cartgrp)

	protected := r.Group("/protected")
	protected.Use(
		middleware.TokenAuthAndRefreshHandler(),
	)
	accountpage.AccountPageRegister(protected)

	addressgrp := protected.Group("/address")
	address.AddressRegister(addressgrp)

	portnum := os.Getenv("APP_PORT")
	if portnum == "" {
		portnum = ":3000"
	}
	fmt.Println(fmt.Sprintf("PortNumber set to %s", portnum))

	r.Run(fmt.Sprintf(":%s", portnum))
}
