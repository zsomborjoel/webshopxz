package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func GetRoot(c *gin.Context) sessions.Session {
	session := sessions.Default(c)
	session.Options(sessions.Options{Path: common.Root})
	return session
}
