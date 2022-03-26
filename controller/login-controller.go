package controller

import (
	"fmt"
	"gin-exercise/entity"
	rest_error "gin-exercise/error"
	"gin-exercise/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context)
	ValidateToken(string) bool
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller loginController) Login(ctx *gin.Context) {
	var credential entity.LoginInfo

	err := ctx.ShouldBind(&credential)
	if err != nil {
		rest_error.NewBadRequestError("No data found")
		return
	}
	isUserAuthenticated := controller.loginService.Login(credential.Email, credential.Password)
	if isUserAuthenticated {

		token := controller.jWtService.GenerateToken(credential.Email)

		if token != "" {

			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}

	}

}

func (controller loginController) ValidateToken(token string) bool {
	newToken, err := controller.jWtService.ValidateToken(token)
	if err != nil {
		rest_error.NewUnAutherizedError("not authorized")

		return false
	}

	if newToken != nil {
		fmt.Println("authorized")

	}
	return true
}
