package middleware

import (
	"errors"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/auth/authtoken"
	"github.com/zsomborjoel/workoutxz/internal/auth/refreshtoken"
	"github.com/zsomborjoel/workoutxz/internal/common"
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

func TokenAuthAndRefreshHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := os.Getenv("JWT_KEY")

		session := sessions.Default(c)
		at := session.Get(common.AccessToken).(string)
		rt := session.Get(common.RefreshToken).(string)

		token, err := jwt.Parse(at, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Jwt Token parse error: %w", err))
		}

		if token.Valid {
			c.Next()
			return
		}

		if refreshtoken.IsValid(rt) {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				c.AbortWithError(http.StatusInternalServerError, errors.New("Jwt claims are not ok"))
				return
			}

			userId := claims["userId"].(string)
			newAccessToken, err := authtoken.CreateJWTToken(userId)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, errors.New("Jwt creation error on refresh"))
				return
			}

			session.Set(common.AccessToken, newAccessToken)
			session.Save()
			c.Next()
			return
		}

		c.AbortWithError(http.StatusUnauthorized, errors.New("Jwt token is Invalid"))
	}
}
