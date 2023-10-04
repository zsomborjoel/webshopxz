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
		c.Next()

		for _, err := range c.Errors {
			log.Error().Err(err).Msg("http error")
		}
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

func CSRFProtectionHandler() gin.HandlerFunc {
	csrfSecret := os.Getenv("CSRF_SECRET")
	if csrfSecret == "" {
		log.Fatal().Msg("CSRF_SECRET environment variable is not set")
	}

	return csrf.Middleware(csrf.Options{
		Secret: csrfSecret,
		ErrorFunc: func(c *gin.Context) {
			c.String(403, "Invalid CSRF token")
			c.Abort()
		},
	})
}

func XSSProtectionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		for key, values := range c.Request.URL.Query() {
			for i, value := range values {
				escapedValue := html.EscapeString(value)
				c.Request.URL.Query()[key][i] = escapedValue
			}
		}
	
		c.Next()
	}
}
