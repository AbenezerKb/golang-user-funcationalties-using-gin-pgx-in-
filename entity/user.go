package entity

type User struct {
	FirstName   string `json:"firstname" binding:"required"`
	SecondName  string `json:"secondname" binding:"required"`
	Age         int    `json:"age" validate:"gte=1, lte=130"`
	Email       string `json:"email" binding:"required,email"`
	Profile     string `json:"profile" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PhoneNumber string `json:"phonenumber" binding:"required"`
}

type LoginInfo struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

// type Item struct {
// 	Title       string `json:"title" binding:"min=2, max=10"`
// 	Description string `json:"description" binding:"max=20"`
// 	URL         string `json:"url" binding:"required,url"`
// 	Owner       Person `json:"owner" binding:"required"`
// }
