package address

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/auth/session"
	"github.com/zsomborjoel/workoutxz/internal/common"
)

func AddressRegister(r *gin.RouterGroup) {
	r.POST("/creation", Creation)
}

func Creation(c *gin.Context) {
	log.Debug().Msg("Creation called")

	co := c.PostForm("country")
	pc := c.PostForm("postal-code")
	ct := c.PostForm("city")
	al := c.PostForm("address-line")
	pn := c.PostForm("phone-number")
	cn := c.PostForm("company-name")
	dt := c.PostForm("details")

	emptyFields := []string{}

	if co == "" {
		emptyFields = append(emptyFields, "Country")
	}
	if pc == "" {
		emptyFields = append(emptyFields, "Postal code")
	}
	if ct == "" {
		emptyFields = append(emptyFields, "City")
	}
	if al == "" {
		emptyFields = append(emptyFields, "Address line")
	}

	if len(emptyFields) > 0 {
		common.AbortWithHtml(c, http.StatusBadRequest, fmt.Sprintf("The following field(s) need to be filled: [%s]", strings.Join(emptyFields, ", ")))
		return
	}

	session := session.GetRoot(c)
	userId := session.Get(common.UserId).(string)

	s := AddressDeserializer{c, AddressRequest{
		Country:     co,
		PortalCode:  pc,
		City:        ct,
		AddressLine: al,
		PhoneNumber: pn,
		CompanyName: cn,
		Details:     dt,
		UserId:      userId,
	}}
	a, err := s.Model()
	if err != nil {
		common.AbortWithHtml(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = UpsertOne(a)
	if err != nil {
		log.Error().Err(err).Msg("")
		common.AbortWithHtml(c, http.StatusInternalServerError, fmt.Sprintf("Internal error occured - try again later"))
		return
	}

	c.Header(common.HTMXRedirect, "/protected/account")
	c.Status(http.StatusOK)
}
