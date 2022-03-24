package controller

import (
	"gin-exercise/entity"
	"gin-exercise/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() []string
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return controller{service: service}
}

func (c controller) FindAll() []string {
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.Save(user)
	return nil
}
