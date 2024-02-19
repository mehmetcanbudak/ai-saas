package util

import (
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,8}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsValidPassword(password string) (string, bool) {
	if len(password) < 8 {
		return "Password is too short", false
	}
	return "Password is valid", true
}
