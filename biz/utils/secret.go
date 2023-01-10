package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/lutasam/GIN_LUTA/biz/common"
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword encrypts a password using bcrypt.
// Cannot decrypt the password! Using ValidatePassword func to validate the password.
func EncryptPassword(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

func ValidatePassword(secret, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(secret), []byte(password))
	if err != nil {
		return common.PASSWORDISERROR
	}
	return nil
}

// AesEncrypt encrypts string using AES encryption algorithm.
// Can decrypt string using AESDecrypt func.
func AesEncrypt(orig string) (string, error) {
	origData := []byte(orig)
	key := []byte(common.OTHERSECRETSALT)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted), nil
}

func AesDecrypt(cryted string) (string, error) {
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	key := []byte(common.OTHERSECRETSALT)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	orig := make([]byte, len(crytedByte))
	blockMode.CryptBlocks(orig, crytedByte)
	orig = PKCS7UnPadding(orig)
	return string(orig), nil
}

func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
