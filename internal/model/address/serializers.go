package address

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

type AddressResponse struct {
	Id          string
	Country     string
	PortalCode  string
	City        string
	AddressLine string
	PhoneNumber string
	CompanyName string
	Details     string
	UserId      string
}

type AddressRequest struct {
	Id          string
	Country     string
	PortalCode  string
	City        string
	AddressLine string
	PhoneNumber string
	CompanyName string
	Details     string
	UserId      string
}

type AddressSerializer struct {
	C *gin.Context
	Address
}

type AddressDeserializer struct {
	C *gin.Context
	AddressRequest
}

func (s *AddressSerializer) Response() AddressResponse {
	return AddressResponse{
		Id:          s.Address.Id,
		Country:     s.Address.Country,
		PortalCode:  s.Address.PortalCode,
		City:        s.Address.City,
		AddressLine: s.Address.AddressLine,
		PhoneNumber: s.Address.PhoneNumber,
		CompanyName: s.Address.CompanyName,
		Details:     s.Address.Details,
		UserId:      s.Address.UserId,
	}
}

func (s *AddressDeserializer) Model() (Address, error) {
	log.Debug().Msg("address.Model called")

	uuid, err := uuid.NewV4()
	if err != nil {
		return Address{}, fmt.Errorf("An error occured in address.Model.NewV4: %w", err)
	}

	return Address{
		Id:          uuid.String(),
		Country:     s.AddressRequest.Country,
		PortalCode:  s.AddressRequest.PortalCode,
		City:        s.AddressRequest.City,
		AddressLine: s.AddressRequest.AddressLine,
		PhoneNumber: s.AddressRequest.PhoneNumber,
		CompanyName: s.AddressRequest.CompanyName,
		Details:     s.AddressRequest.Details,
		UserId:      s.AddressRequest.UserId,
	}, nil
}
