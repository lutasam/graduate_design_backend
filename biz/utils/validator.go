package utils

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

func IsValidURL(url string) bool {
	return govalidator.IsURL(url)
}

func IsValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}

func IsValidPhoneNumber(phoneNumber string) bool {
	regRuler := "^1[345789]{1}\\d{9}$" // chinese phone number rule
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(phoneNumber)
}
