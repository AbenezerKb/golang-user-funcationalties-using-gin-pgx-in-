package main

import (
	"gin-exercise/controller"
	"gin-exercise/middlewares"
	"gin-exercise/service"
	"io"
	"os"

	//ginDumb "github.com/tpkeeper/gin-dumb"

	//"middlewares"
	"github.com/gin-gonic/gin"
)

var (
	services service.UserService       = service.New()
	controll controller.UserController = controller.New(services)
	//kk controller.LoginController
	loginService    service.LoginService       = &service.UserLogin{}
	jwtService      service.JWTService         = service.JWTAuthService()
	loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)
	//loginController :=controller.LoginController {}
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

	//get users list
	server.GET("/userslist", controll.FindAll)

	//add new user
	server.POST("/users", controll.Save)

	//login
	server.POST("/login", loginController.Login)

	server.Run(":8080")
}
