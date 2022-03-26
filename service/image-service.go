package service

import (
	//"fmt"
	"net/http"
	//"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ImageService interface {
	Save(ctx *gin.Context)    //(*entity.User, *rest_error.RestErr)
	Display(ctx *gin.Context) //string
}

type imageService struct {
	image string
}

func NewImage() ImageService {
	return &imageService{}
}

func (i *imageService) Save(ctx *gin.Context) /*(*entity.User, *rest_error.RestErr)*/ {

	saveFileHandler(ctx)
	// //TODO handle it using gin
	// in, fileHeader, err := ctx.FormFile("image")
	// if err != nil {
	// 	rest_error.NewBadRequestError("image upload failed")
	// }
	// //fileHeader.Filename
	// //defer in.Close()

	// // filename, err := ctx.FormFile("image")
	// // if err !=nil{
	// // 	rest_error.NewBadRequestError("image upload failed")
	// // }
	// defer in.Close()

	// //path.Join("./static",image)
	// out, err := os.OpenFile(fileHeader.Filename, os.O_WRONLY, 0644)
	// if err != nil {
	// 	rest_error.NewInternalServerError("image saving failed")
	// }
	// defer out.Close()
	// io.Copy(out, in)
	// pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	// if err != nil {
	// 	return nil, rest_error.NewInternalServerError("password encryption falied")
	// }
	// user.Password = string(pw[:])
	// fmt.Println(user)
	// db.SaveUser(user)
	//	return &user, nil
}

func (i *imageService) Display(ctx *gin.Context) {
	// fn := ctx.Param("userid")
	// str, _ := os.Stat(filepath.Join("v1/image/", fn, ".jpg"))
	// static.Serve("/img", static.LocalFile("./img", true))
	// ctx.Static("/img", "./img")
	// ctx.
	// ctx.File(str.Name())

	// fmt.Println("herer,", filepath.Join("v1/image/", fn))
	// if _, err := os.Stat(filepath.Join("v1/image/", fn, ".jpg")); err == nil {
	// 	fmt.Println("file exists and returned", filepath.Join("v1/image/", fn))
	// 	ctx.JSON(200, filepath.Join("v1/image/", fn))
	// }

	//return "db.Userslist()"
}

func saveFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	fn := c.Param("userid")
	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	newFileName := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFile := fn + newFileName

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "C:\\Users\\Administrator\\Documents\\Gin-exercises\\v1\\image\\"+newFile); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}
