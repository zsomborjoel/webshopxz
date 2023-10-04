package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/auth/authtoken"
	"github.com/zsomborjoel/workoutxz/internal/auth/refreshtoken"
	"github.com/zsomborjoel/workoutxz/internal/auth/verificationtoken"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"github.com/zsomborjoel/workoutxz/internal/model/user"
)

func AuthRegister(r *gin.RouterGroup) {
	r.POST("/registration", Registration)
	r.GET(common.ConfirmRegistrationEndpoint, ConfirmRegistration)
	r.PUT("/resend-verification", ResendVerification)
	r.POST("/login", Login)
	r.POST("/refresh-token", RefreshJWTToken)
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
		common.AbortWithHtml(c, http.StatusBadRequest, "Email is not valid")
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

	rr := RegistrationRequest{e, e, p}

	var u user.User
	var err error
	s := RegistrationRequestSerializer{c, rr}
	u, err = s.Model()
	if err != nil {
		common.AbortWithHtml(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = user.ExistByUserName(u.UserName)
	if err != nil {
		common.AbortWithHtml(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := user.CreateOne(u); err != nil {
		common.AbortWithHtml(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = verificationtoken.CreateOne(u)
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
	lr := &LoginRequest{}
	if err := c.BindJSON(&lr); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid login request body"))
		return
	}

	usr, err := user.FindByUserName(lr.UserName)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	err = validatePassword(lr.Password, usr.Password)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	jwt, err := authtoken.CreateJWTToken(usr)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rt, err := refreshtoken.CreateOne(usr)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	s := JwtTokenSerializer{c, usr, jwt, rt}
	c.JSON(http.StatusOK, s.Response())
}

func RefreshJWTToken(c *gin.Context) {
	log.Debug().Msg("RefreshJWTToken called")

	rcvt := c.Query("refreshtoken")

	if rcvt == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid refreshtoken"))
		return
	}

	rt, err := refreshtoken.IsValid(rcvt)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	usr, err := user.FindByUserId(rt.UserId)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	jwt, err := authtoken.CreateJWTToken(usr)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	s := RefreshTokenSerializer{c, jwt, rt.Token}
	c.JSON(http.StatusOK, s.Response())
}
