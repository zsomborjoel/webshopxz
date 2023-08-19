package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/auth"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var finalErr error
		for _, err := range c.Errors {
			finalErr = err
			log.Error().Err(err).Msg("http error")
		}

		// status -1 doesn't overwrite existing status code
		c.JSON(-1, finalErr)
	}
}

func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		key := os.Getenv("JWT_KEY")
		bearer := c.GetHeader("Authorization")

		if !strings.HasPrefix(bearer, "Bearer") {
			c.AbortWithError(http.StatusUnauthorized, errors.New("It is not a Bearer token"))
		}
		jwtToken := bearer[7:]

		token, err := jwt.ParseWithClaims(jwtToken, &auth.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Jwt Token parse error: %w", err))
		}

		if !token.Valid {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Jwt token is Invalid"))
		}
	}
}
