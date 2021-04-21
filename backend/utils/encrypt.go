package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SecurityHashing(original string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(original), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hash), nil
}

func ValidateSecurityHashing(origin, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(origin))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
