package user

import "github.com/gin-gonic/gin"

type UserResponse struct {
	Id       string
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type UserSerializer struct {
	C *gin.Context
	User
}

func (s *UserSerializer) Response() UserResponse {
	return UserResponse{
		Id:       s.User.Id,
		UserName: s.User.UserName,
		Email:    s.User.Email,
	}
}
