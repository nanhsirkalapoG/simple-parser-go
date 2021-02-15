package utils

import (
	"fmt"
	"parser/app/constants"
)

// Used regex to validate the email
// To validate email by sending it - https://github.com/badoux/checkmail
// Worth read: https://davidcel.is/posts/stop-validating-email-addresses-with-regex/
func ValidEmail(email string) bool  {
	return constants.EmailRegex.MatchString(email)
}

// Validate the employee id
func ValidId(id string) bool  {
	return true
}

// Validate the employee firstName
func ValidFirstName(firstName string) bool  {
	return true
}

// Validate the employee lastName
func ValidLastName(lastName string) bool  {
	return true
}

// Validate the employee fullName
func ValidFullName(fullName string) bool  {
	return true
}

// Validate the employee salary
func ValidSalary(salary string) bool  {
	if _, err := GetFormattedSalary(salary); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
