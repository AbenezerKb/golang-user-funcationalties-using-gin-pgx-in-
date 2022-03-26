package controller

import (
	"gin-exercise/service"

	//"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin"
	//    "github.com/google/uuid" // To generate random file names
)

type ImageController interface {
	ImageSave(ctx *gin.Context)
	Display_image(ctx *gin.Context)
}

type imgcontroller struct {
	service service.ImageService
}

func NewImage(service service.ImageService) ImageController {
	return imgcontroller{service: service}
}

func (c imgcontroller) ImageSave(ctx *gin.Context) {

	//ctx.JSON(200, c.service.Save(ctx))
	c.service.Save(ctx)

}

func (c imgcontroller) Display_image(ctx *gin.Context) {

	c.service.Display(ctx)

	// var user entity.User

	// err := ctx.ShouldBindJSON(&user)
	// if err != nil {
	// 	rest_error.NewBadRequestError(("error, registration failed"))
	// }
	// c.service.Save(ctx)
	// ctx.JSON(http.StatusOK, gin.H{"Message": "user input is valid!"})

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "Your file has been successfully uploaded.",
	// })

}
