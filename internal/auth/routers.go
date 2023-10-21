package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/auth/refreshtoken"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	authtoken "github.com/zsomborjoel/workoutxz/internal/auth/token"
	"github.com/zsomborjoel/workoutxz/internal/auth/verificationtoken"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/model/user"
)

func AuthRegister(r *gin.RouterGroup) {
	r.POST("/registration", Registration)
	r.GET(common.ConfirmRegistrationEndpoint, ConfirmRegistration)
	r.PUT("/resend-verification", ResendVerification)
	r.POST("/login", Login)
	r.POST("/logout", Logout)
	r.POST("/reset-password", ResetPassword)
}

func Registration(c *gin.Context) {
	log.Debug().Msg("Registration called")

	e := c.PostForm("email")
	p := c.PostForm("password")
	cp := c.PostForm("confirm-password")
	t := c.PostForm("terms")

	if e == "" || p == "" {
		common.AbortWithHtml(c, http.StatusBadRequest, "Email or password can not be empty")
		return
	}

	if !common.IsValidEmail(e) {
		common.AbortWithHtml(c, http.StatusBadRequest, "Email is not in valid format")
		return
	}

	if p != cp {
		common.AbortWithHtml(c, http.StatusBadRequest, "Password and Confirm Password is not equal")
		return
	}

	if t != "on" {
		common.AbortWithHtml(c, http.StatusBadRequest, "Please approve Terms and Conditions")
		return
	}

	var usr user.User
	var err error
	rr := RegistrationRequest{e, e, p}
	s := RegistrationRequestSerializer{c, rr}
	usr, err = s.Model()
	if err != nil {
		common.AbortWithHtml(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user.ExistByUserName(usr.UserName) {
		common.AbortWithHtml(c, http.StatusBadRequest, "User alredy exists")
		return
	}

	if err := user.CreateOne(usr); err != nil {
		common.AbortWithHtml(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = verificationtoken.CreateOne(usr)
	if err != nil {
		common.AbortWithHtml(c, http.StatusInternalServerError, err.Error())
		return
	}

	/*if err := email.SendEmail(u.Email, t); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}*/

	common.OkWithHtml(c, "Account been created")
}

func ConfirmRegistration(c *gin.Context) {
	log.Debug().Msg("ConfirmRegistration called")

	t := c.Query("token")

	if t == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid token"))
		return
	}

	vt, err := verificationtoken.IsValid(t)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	err = user.ActivateOne(user.User{Id: vt.UserId})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	verificationtoken.DeleteOne(t)

	c.Writer.WriteHeader(http.StatusOK)
}

func ResendVerification(c *gin.Context) {
	log.Debug().Msg("ResendVerification called")

	t := c.Query("verificationtoken")

	if t == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid verificationtoken"))
		return
	}

	_, err := verificationtoken.IsValid(t)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	new, err := verificationtoken.UpdateToken(t)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, new)
}

func Login(c *gin.Context) {
	log.Debug().Msg("Login called")

	e := c.PostForm("email")
	p := c.PostForm("password")

	usr, err := user.FindByUserName(e)
	if err != nil {
		log.Debug().Err(err).Msg(common.LoginError)
		common.AbortWithHtml(c, http.StatusNotFound, "Username or password is invalid")
		return
	}

	err = validatePassword(usr.Password, p)
	if err != nil {
		log.Debug().Err(err).Msg(common.LoginError)
		common.AbortWithHtml(c, http.StatusUnauthorized, "Username or password is invalid")
		return
	}

	jwt, err := authtoken.CreateJWTToken(usr.Id)
	if err != nil {
		log.Debug().Err(err).Msg(common.LoginError)
		common.AbortWithHtml(c, http.StatusInternalServerError, "Internal server error on login")
		return
	}

	rt, err := refreshtoken.CreateOne(usr)
	if err != nil {
		log.Debug().Err(err).Msg(common.LoginError)
		common.AbortWithHtml(c, http.StatusInternalServerError, "Internal server error on login")
		return
	}

	session := session.GetRoot(c)
	session.Set(common.AccessToken, jwt)
	session.Set(common.RefreshToken, rt)
	session.Set(common.UserId, usr.Id)
	session.Save()

	c.Header(common.HTMXRedirect, "/")
	c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {
	log.Debug().Msg("Logout called")

	session := session.GetRoot(c)
	session.Clear()
	session.Save()

	c.Header(common.HTMXRedirect, "/")
	c.Status(http.StatusOK)
}

func ResetPassword(c *gin.Context) {
	log.Debug().Msg("ResetPassword called")

	e := c.PostForm("email")

	if e == "" {
		common.AbortWithHtml(c, http.StatusBadRequest, "Email can not be empty")
		return
	}

	if !common.IsValidEmail(e) {
		common.AbortWithHtml(c, http.StatusBadRequest, "Email is not in valid format")
		return
	}

	if !user.ExistByUserName(e) {
		common.AbortWithHtml(c, http.StatusBadRequest, "Email not exists in our system")
		return
	}

	err := user.ResetPasswordByUserName(e)
	if err != nil {
		common.AbortWithHtml(c, http.StatusInternalServerError, "We were not able to reset your password - contact us via email")
		return
	}

	common.OkWithHtml(c, "Password been reseted - Check for the verification in your email account!")
}
