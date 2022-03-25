package controller

import (
	"gin-exercise/entity"
	rest_error "gin-exercise/error"
	"gin-exercise/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return controller{service: service}
}

func (c controller) FindAll(ctx *gin.Context) {

	ctx.JSON(200, c.service.FindAll())

}

func (c controller) Save(ctx *gin.Context) {
	var user entity.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		rest_error.NewBadRequestError(("error, registration failed"))
	}
	c.service.Save(user)
	ctx.JSON(http.StatusOK, gin.H{"Message": "user input is valid!"})
}
