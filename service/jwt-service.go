package service

import (
	"fmt"
	rest_error "gin-exercise/error"
	"time"

	"github.com/golang-jwt/jwt"
)

//jwt service
type JWTService interface {
	GenerateToken(email string) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Email string `json:"email"`
	//	Uuu   bool   `json:"uuu"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	//issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: "mysecretkey",
		//	issure:    "Abenezer",
	}
}

// func getSecretKey() string {
// 	secret := os.Getenv("KEY")
// 	if secret == "" {
// 		secret = "mysecretkey"
// 	}
// 	return secret
// }

func (service *jwtServices) GenerateToken(email string) string {
	claims := &authCustomClaims{
		email,
		//	uuu,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		rest_error.NewInternalServerError("secret key error")
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
