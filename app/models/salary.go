package models

import (
	"errors"
	"parser/app/constants"
	"parser/app/utils"
	"strconv"
)

type Salary struct {
	Wage string
	FieldName string
}

// Validate the employee salary and set it to the object
func (s *Salary) SetValue(record []string, headerMap map[string]int) error {
	fieldIndex := utils.GetMatchedIndex(constants.SalaryVariations, headerMap)
	if fieldIndex == -1 {
		return errors.New("employee salary is missing")
	}
	if record[fieldIndex] == "" || !utils.ValidSalary(record[fieldIndex]) {
		return errors.New("invalid employee salary")
	}

	formattedWage, _ := utils.GetFormattedSalary(record[fieldIndex])
	s.Wage = strconv.FormatFloat(formattedWage, 'f', -1, 64)
	return nil
}

// Get the salary from the object
func (s Salary) GetValue() interface{} {
	return s.Wage
}

// Get the field name from the object
func (s Salary) GetFieldName() string {
	return s.FieldName
}