package middleware

import (
	"errors"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/auth/authtoken"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var finalErr error
		for _, err := range c.Errors {
			finalErr = err
			log.Error().Err(err).Msg("http error")
		}

		if finalErr != nil {
			c.JSON(-1, finalErr)
		}

		c.Next()
	}
}

func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := os.Getenv("JWT_KEY")
		bearer := c.GetHeader("Authorization")

		if !strings.HasPrefix(bearer, "Bearer") {
			c.AbortWithError(http.StatusUnauthorized, errors.New("It is not a Bearer token"))
		}
		jwtToken := bearer[7:]

		token, err := jwt.ParseWithClaims(jwtToken, &authtoken.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Jwt Token parse error: %w", err))
		}

		if !token.Valid {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Jwt token is Invalid"))
		}

		c.Next()
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

func XSSProtectionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		writer := c.Writer
		customWriter := &responseWriterWithInterceptor{ResponseWriter: writer}

		c.Writer = customWriter

		responseBody := customWriter.Body()
		sanitizedResponse := html.EscapeString(responseBody)

		writer.Write([]byte(sanitizedResponse))
		c.Next()
	}
}

func CSRFProtectionHandler() gin.HandlerFunc {
	csrfkey := os.Getenv("CSRF_KEY")
	if csrfkey == "" {
		log.Error().Msg("CSRF_KEY environment variable is not set")
	}

	return csrf.Middleware(csrf.Options{
		Secret: csrfkey,
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	})
}
