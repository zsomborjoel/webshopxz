package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			log.Error().Err(err).Msg("http error")
		}
	}
}

// Needed for static files like images
func StaticFileHandler() gin.HandlerFunc {
	root := os.Getenv("STATIC_FILE_PATH")
	if root == "" {
		log.Error().Msg("STATIC_FILE_PATH environment variable is not set")
	}

	prefix := os.Getenv("STATIC_FILE_PREFIX")
	if prefix == "" {
		log.Error().Msg("STATIC_FILE_PREFIX environment variable is not set")
	}

	fileServer := http.FileServer(http.Dir(root))

	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, prefix) {
			c.Next()
			return
		}

		c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, prefix)
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
