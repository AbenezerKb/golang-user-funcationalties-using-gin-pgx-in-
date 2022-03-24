package service

import (
	"gin-exercise/db"
	"gin-exercise/entity"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []string
}

type userService struct {
	users []entity.User
}

func New() UserService {
	return &userService{}
}

func (i *userService) Save(user entity.User) entity.User {
	i.users = append(i.users, user)
	db.SaveUser(user)
	return user
}

func (i *userService) FindAll() []string {

	return db.Userslist()
}
