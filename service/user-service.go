package service

import (
	"fmt"
	"gin-exercise/db"
	"gin-exercise/entity"
	rest_error "gin-exercise/error"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Save(entity.User) (*entity.User, *rest_error.RestErr)
	FindAll() []string
}

type userService struct {
	users []entity.User
}

func New() UserService {
	return &userService{}
}

func (i *userService) Save(user entity.User) (*entity.User, *rest_error.RestErr) {
	pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, rest_error.NewBadRequestError("password encryption falied")
	}
	user.Password = string(pw[:])
	fmt.Println(user)
	db.SaveUser(user)
	return &user, nil
}

func (i *userService) FindAll() []string {

	return db.Userslist()
}
