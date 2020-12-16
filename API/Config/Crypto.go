package Config

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashGen(str string) string{
	pwd := []byte(str) //Convert string to byte
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash) //return string of byte
}

func HashCompare(pwd,hashed string) bool{
	byteHash := []byte(hashed)
	bytePw   := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePw)
	if err != nil {
		return false
	}else {
		return true
	}
}