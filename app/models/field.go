package models

/*
	This class is to manage the different attributes of employee structure
 */
import (
	"parser/app/constants"
	"reflect"
)

type iFields interface {
	SetValue(record []string, headerMap map[string]int) error
	GetValue() interface{}
	GetFieldName() string
}

type Fields struct {
	fieldInterface iFields
}

var fieldList = []iFields{
	&Email{FieldName:"Email"},
	&Id{FieldName:"Id"},
	&Name{FieldName:"Name"},
	&Salary{FieldName:"Salary"},
}

// Set the different field values to from employee
func (f Fields) FormData(record []string, headerMap map[string]int) error {
	for _, field := range fieldList {
		err := field.SetValue(record, headerMap)
		if err != nil {
			return err
		}
	}
	return nil
}

// Get the data to from different attributes and form employee structure
func (f Fields) GetData() employee {
	emp := employee{}
	for _, field := range fieldList {
		switch (field.GetValue()).(type) {
		case string:
			reflect.ValueOf(&emp).Elem().FieldByName(field.GetFieldName()).SetString(field.GetValue().(string))
		}
	}

	return emp
}

// Headers for valid CSV file
func (f Fields) GetHeader() []string {
	return constants.Headers
}