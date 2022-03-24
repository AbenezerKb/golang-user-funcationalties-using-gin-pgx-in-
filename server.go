package main

import (
	"fmt"
	"gin-exercise/controller"
	"gin-exercise/middlewares"
	"gin-exercise/service"
	"io"
	"net/http"
	"os"

	//ginDumb "github.com/tpkeeper/gin-dumb"

	//"middlewares"
	"github.com/gin-gonic/gin"
)

var (
	services service.UserService       = service.New()
	controll controller.UserController = controller.New(services)
	//loginService    service.LoginService       = service.StaticLoginService()
	//jwtService      service.JWTService         = service.JWTAuthService()
	//loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)
)

func setUptLogOutPut() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	//gin.DefaultWriter := io.MultiWriter(file,os.Stdout)

	//gin.DefaultWriter := io.MultiWriter(file, os.Stdout)
}

func main() {
	//setUptLogOutPut()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger()) //, ginDumb.Dumb() , middlewares.AuthorizeJWT()
	fmt.Println("working till here 1")
	//get users list
	server.GET("/userslist", func(ctx *gin.Context) {
		fmt.Println("working till here 2")
		ctx.JSON(200, controll.FindAll())
		fmt.Println("working till here 3")
	})

	//add new user
	server.POST("/users", func(ctx *gin.Context) {
		err := controll.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"Message": "user input is valid!"})
		}

		ctx.JSON(200, controll.Save(ctx))
	})

	//login
	server.POST("/login", func(ctx *gin.Context) {
		token := "loginController.Login(ctx)"
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}

		ctx.JSON(200, controll.Save(ctx))
	})

	server.Run(":8080")
}
