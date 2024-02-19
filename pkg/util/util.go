package util

import (
	"regexp"
	"strings"
	"unicode"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,8}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// validatePassword checks if the password is strong and meets the criteria:
// - At least 8 characters long
// - Contains at least one digit
// - Contains at least one lowercase letter
// - Contains at least one uppercase letter
// - Contains at least one special character
// func IsValidPassword(password string) (string, bool) {
func IsValidPassword(password string) (string, bool) {

	var (
		hasUpper     = false
		hasLower     = false
		hasNumber    = false
		hasSpecial   = false
		specialRunes = "!@#$%^&*"
	)

	if len(password) < 8 {
		return "Password must be at least 8 characters long", false
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char) || strings.ContainsRune(specialRunes, char):
			hasSpecial = true
		}
	}

	//Return true only if all conditions are met
	return "Password must be at least 8 characters long, containss at least one digit-one lowercase letter-one uppercase letter and one special character.", hasUpper && hasLower && hasNumber && hasSpecial
	//return "", hasUpper && hasLower && hasNumber && hasSpecial
	//return hasUpper && hasLower && hasNumber && hasSpecial

}
