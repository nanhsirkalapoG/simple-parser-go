package models

import (
	"errors"
	"fmt"
	"go/types"
	"parser/app/constants"
	"parser/app/utils"
	"strings"
)

type Name struct {
	FirstName string
	LastName  string
	FullName  string
	FieldName string
}

type iName interface {
	GetName() string
	SetName(name Name) types.Nil
}

// Validate the employee name and set it to the object
func (n *Name) SetValue(record []string, headerMap map[string]int) error {
	n.setFirstName(record, headerMap)
	n.setLastName(record, headerMap)
	return n.setFullName(record, headerMap)
}

// Validate and set first name
func (n *Name) setFirstName(record []string, headerMap map[string]int) error {
	fieldIndex := utils.GetMatchedIndex(constants.FirstNameVariations, headerMap)
	if fieldIndex != -1 {
		n.FirstName = strings.TrimSpace(record[fieldIndex])
	}
	return nil
}

// Validate and set last name
func (n *Name) setLastName(record []string, headerMap map[string]int) error {
	fieldIndex := utils.GetMatchedIndex(constants.LastNameVariations, headerMap)
	if fieldIndex != -1 {
		n.LastName = strings.TrimSpace(record[fieldIndex])
	}
	return nil
}

// Validate and set full name
func (n *Name) setFullName(record []string, headerMap map[string]int) error {
	fieldIndex := utils.GetMatchedIndex(constants.FullNameVariations, headerMap)
	if fieldIndex == -1 {
		if n.FirstName == "" || n.LastName == "" {
			return errors.New("employee name is missing")
		}
		n.FullName = fmt.Sprintf("%s %s", n.FirstName, n.LastName)
	} else {
		n.FullName = strings.TrimSpace(record[fieldIndex])
	}
	if n.FullName == "" || !utils.ValidId(n.FullName) {
		return errors.New("invalid employee name")
	}

	return nil
}

// Get the full name from the object
func (n Name) GetValue() interface{} {
	return n.FullName
}

// Get the field name from the object
func (n Name) GetFieldName() string {
	return n.FieldName
}
