package main

import (
	"gin-exercise/controller"
	"gin-exercise/middlewares"
	"gin-exercise/service"
	"path"

	//ginDumb "github.com/tpkeeper/gin-dumb"

	//"middlewares"
	"github.com/gin-gonic/gin"
)

var (
	services service.UserService       = service.New()
	controll controller.UserController = controller.New(services)

	loginService    service.LoginService       = &service.UserLogin{}
	jwtService      service.JWTService         = service.JWTAuthService()
	loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	imgservices service.ImageService       = service.NewImage()
	imgcontroll controller.ImageController = controller.NewImage(imgservices)
)

func main() {
	//setUptLogOutPut()
	server := gin.New()

	server.Static("/v1/image", "./")

	server.Use(gin.Recovery(), middlewares.Logger()) //, ginDumb.Dumb() , middlewares.AuthorizeJWT()

	//get users list
	server.GET("/users", controll.FindAll)

	//add new user
	server.POST("/users", controll.Save)

	//login
	server.POST("/login", loginController.Login)

	//image upload
	server.POST("/user/:userid/upload_image", imgcontroll.ImageSave)

	//access image
	//server.GET("/image/:userid", imgcontroll.Display_image)
	server.Static("/image", path.Join("v1", "image"))
	server.Run(":8080")
}
