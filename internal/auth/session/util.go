package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"github.com/zsomborjoel/workoutxz/internal/common/consts"
)

func GetRoot(c *gin.Context) sessions.Session {
	session := sessions.Default(c)
	session.Options(sessions.Options{Path: consts.Root})
	return session
}

func SetCsrfTokenCookie(c *gin.Context) {
	c.SetCookie("csrf_token", csrf.GetToken(c), 0, "/", "", false, false)
}
