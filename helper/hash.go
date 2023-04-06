package helper

import (
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(pass string) (string, error) {
	raw, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(raw), nil
}

func ComparePassword(hashed string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}
