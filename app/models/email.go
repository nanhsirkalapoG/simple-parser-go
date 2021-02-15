package models

import (
	"errors"
	"parser/app/constants"
	"parser/app/utils"
	"strings"
)

type Email struct {
	Id     string
	FieldName string
}

// Validate the email and set it to the object
func (e *Email) SetValue(record []string, headerMap map[string]int) error {
	fieldIndex := utils.GetMatchedIndex(constants.EmailVariations, headerMap)
	if fieldIndex == -1 {
		return errors.New("email id is missing")
	}

	if record[fieldIndex] == "" || !utils.ValidEmail(record[fieldIndex]) {
		return errors.New("invalid email address")
	}
	e.Id = strings.TrimSpace(strings.ToLower(record[fieldIndex]))
	return nil
}

// Get the email id from the object
func (e Email) GetValue() interface{} {
	return e.Id
}

// Get the field name from the object
func (e Email) GetFieldName() string {
	return e.FieldName
}
