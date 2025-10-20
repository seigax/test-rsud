package lib

import (
	"log"
	"net/mail"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CustomFunction() (string, error) {
	return "", nil
}

func GenerateHashFromString(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	return err == nil
}

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsIn[T comparable](array []T, data T) bool {
	for i := range array {
		if array[i] == data {
			return true
		}
	}

	return false
}

func CreateFolder(name string) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			err = os.Mkdir(name, 0755)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// other error
			log.Fatal(err)
		}
	}
}

func CreateGormLog() (*os.File, error) {
	CreateFolder("logs")
	return os.Create("logs/output-gorm_" + time.Now().Format("2006-01-02H15M04") + ".log")
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
