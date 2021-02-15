package models

import (
	"errors"
	"parser/app/constants"
	"parser/app/utils"
	"strings"
)

type Id struct {
	Identifier string
	FieldName string
}

// Validate the employee id and set it to the object
func (i *Id) SetValue(record []string, headerMap map[string]int) error {
	fieldIndex := utils.GetMatchedIndex(constants.IdVariations, headerMap)
	if fieldIndex == -1 {
		return errors.New("employee id is missing")
	}
	if record[fieldIndex] == "" || !utils.ValidId(record[fieldIndex]) {
		return errors.New("invalid employee id")
	}
	i.Identifier = strings.TrimSpace(record[fieldIndex])
	return nil
}

// Get the employee id from the object
func (i Id) GetValue() interface{} {
	return i.Identifier
}

// Get the field name from the object
func (i Id) GetFieldName() string {
	return i.FieldName
}
