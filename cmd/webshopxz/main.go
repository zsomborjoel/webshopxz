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
	"github.com/zsomborjoel/workoutxz/internal/ping"
	"github.com/zsomborjoel/workoutxz/internal/webpage/mainpage"
)

func main() {
	fmt.Println("Application Init started")
	
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
		middleware.XSSProtectionHandler(),
		middleware.StaticFileHandler(),
		middleware.ErrorHandler(),
	)

	v1 := r.Group("")

	// technical
	ping.PingRegister(v1.Group("/ping"))
	auth.AuthRegister(v1.Group("/auth"))
	email.EmailRegister(v1.Group("/email"))

	// template
	mainpage.MainPageRegister(v1.Group("/mainpage"))

	r.Run()
}
