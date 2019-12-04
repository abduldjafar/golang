package Account

import (
	"backend-code/model"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetToken(email string) string {
	expiresAt := time.Now().Add(time.Hour * 24).Unix()
	tk := &model.Token{
		Email: email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	return tokenString
}
