package controller

import (
	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/internal/services"
	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	// Add fields if necessary
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserName(c *gin.Context) {

	name, error := uc.userService.GetUserName()

	if error != nil {
		response.Error(c, 20003)
	}

	response.Success(c, 20001, name)
}
