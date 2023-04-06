package helper

import (
	"deploy/config"
	"log"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id string) string {
	var informasi = jwt.MapClaims{}
	informasi["id"] = id

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, informasi)

	resultToken, err := rawToken.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		log.Println("generate jwt error ", err.Error())
		return ""
	}

	return resultToken
}

func DecodeJWT(token *jwt.Token) string {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		user_id := data["id"].(string)

		return user_id
	}

	return ""
}
