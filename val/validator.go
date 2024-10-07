package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	if len(value) < minLength {
		return fmt.Errorf("value length is less than %d", minLength)
	}
	if len(value) > maxLength {
		return fmt.Errorf("value length is more than %d", maxLength)
	}
	return nil

}

func ValidateUsername(username string) error {
	if err := ValidateString(username, 3, 50); err != nil {
		return err
	}
	if !isValidUsername(username) {
		return fmt.Errorf("username is invalid")
	}
	return nil

}

func ValidatePassword(password string) error {
	return ValidateString(password, 6, 50)
}

func ValidateEmail(email string) error {
	if err := ValidateString(email, 3, 100); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("email is invalid")
	}

	return nil

}

func ValidateFullName(fullName string) error {
	if err := ValidateString(fullName, 3, 50); err != nil {
		return err
	}

	if !isValidFullName(fullName) {
		return fmt.Errorf("full name is invalid")
	}

	return nil

}
