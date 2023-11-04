package middleware

import (
	"errors"
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/auth/refreshtoken"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	authtoken "github.com/zsomborjoel/workoutxz/internal/auth/token"
	"github.com/zsomborjoel/workoutxz/internal/common/consts"
	"github.com/zsomborjoel/workoutxz/internal/common/response"
)

func CSRFProtectionHandler() gin.HandlerFunc {
	csrfSecret := os.Getenv("CSRF_SECRET")
	if csrfSecret == "" {
		log.Fatal().Msg("CSRF_SECRET environment variable is not set")
	}

	return csrf.Middleware(csrf.Options{
		Secret: csrfSecret,
		ErrorFunc: func(c *gin.Context) {
			msg := "Invalid CSRF token"
			log.Error().Msg(msg)
			c.String(403, msg)
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
		session := session.GetRoot(c)
		at := session.Get(consts.AccessToken)
		rt := session.Get(consts.RefreshToken)

		if at == nil || rt == nil {
			log.Err(errors.New("Token not present")).Msg("")
			response.AbortWithUnauthorizedHtml(c)
			return
		}

		token, err := authtoken.Parse(at.(string))
		if err != nil {
			resetLogin(session)
			log.Err(fmt.Errorf("Jwt Token parse error: %w", err)).Msg("")
			response.AbortWithUnauthorizedHtml(c)
			return
		}

		if token.Valid {
			c.Next()
			return
		}

		if isValidRefreshToken(rt.(string)) {
			handleValidRefreshToken(c, token, session)
			return
		}

		c.AbortWithError(http.StatusUnauthorized, errors.New("Jwt token is Invalid"))
	}
}

func resetLogin(session sessions.Session) {
	session.Delete(consts.AccessToken)
	session.Delete(consts.RefreshToken)
	err := session.Save()
	if err != nil {
		log.Error().Err(err).Msg("Failed to save session security.resetLogin")
	}
}

func isValidRefreshToken(tokenString string) bool {
	return refreshtoken.IsValid(tokenString)
}

func handleValidRefreshToken(c *gin.Context, token *jwt.Token, session sessions.Session) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Jwt claims are not ok"))
		return
	}

	userId := claims["UserId"]
	if userId == "" {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Claims does not contain userId"))
		return
	}

	handleTokenRefresh(c, userId.(string), session)
}

func handleTokenRefresh(c *gin.Context, userID string, session sessions.Session) {
	newAccessToken, err := authtoken.CreateJWTToken(userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Jwt creation error on refresh"))
		return
	}

	session.Set(consts.AccessToken, newAccessToken)
	err = session.Save()
	if err != nil {
		log.Error().Err(err).Msg("Failed to save session security.handleTokenRefresh")
	}

	c.Next()
}
