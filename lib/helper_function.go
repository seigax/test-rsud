package lib

import (
	"crypto/aes"
	"crypto/cipher"
	cryptoRand "crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"math/rand"
	"time"
)

const otpCharset = "0123456789"

func GenerateOTP(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	otp := make([]byte, length)
	for i := range otp {
		otp[i] = otpCharset[generator.Intn(len(otpCharset))]
	}

	return string(otp)
}

func GenerateFutureTimeSeconds(seconds int) time.Time {
	now := time.Now()
	future := now.Add(time.Duration(seconds) * time.Second)
	return future
}

func EncryptKey(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(cryptoRand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func DecryptKey(key []byte, ciphertext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ct, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(ct) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ct[:aes.BlockSize]
	ct = ct[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ct, ct)

	return string(ct), nil
}
