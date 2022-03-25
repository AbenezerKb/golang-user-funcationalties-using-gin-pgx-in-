package service

import (
	"gin-exercise/db"
)

type LoginService interface {
	Login(email string, password string) bool
}

func NewLogin() LoginService {
	return &UserLogin{}
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *UserLogin) Login(email string, password string) bool {
	return db.UserInfo(email, password)
}
