package processors

import (
	"fmt"
	"parser/app/models"
)

type Result struct {
	ValidRecords [][]string
	ValidRecordsHeader []string
	InValidRecords [][]string
	InValidRecordsHeader []string
}

// Process all the file records and form Result structure
func ProcessRecords(records [][]string, headers []string, headerMap map[string]int) (Result, error) {
	result := Result{}
	field := models.Fields{}
	validRecords := [][]string{}
	inValidRecords := [][]string{}
	emailMap := map[string]string{}
	for _, record := range records {
		err := field.FormData(record, headerMap)
		if err != nil {
			inValidRecords = append(inValidRecords, record)
			fmt.Println(fmt.Sprintf("error processing record: %v, error: %v", record, err))
			continue
		}
		employee := field.GetData()
		if emailMap[employee.Email] == "" {
			validRecords = append(validRecords, employee.ToSlice())
			emailMap[employee.Email] = employee.Id
			continue
		}
		inValidRecords = append(inValidRecords, record)
		fmt.Println(fmt.Sprintf("error processing record: %v, error: %v", record, "email already exist"))
	}
	result.InValidRecords = inValidRecords
	result.InValidRecordsHeader = headers
	result.ValidRecords = validRecords
	result.ValidRecordsHeader = field.GetHeader()
	return result, nil
}