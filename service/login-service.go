package service

type LoginService interface {
	Login(username string, password string) bool
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *UserLogin) Login(username string, password string) bool {
	return u.Username == username &&
		u.Password == password
}
