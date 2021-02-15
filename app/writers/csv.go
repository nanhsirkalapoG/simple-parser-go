package writers

import (
	"encoding/csv"
	"os"
)

// Write the give data into a CSV file with headers
func WriteFile(fileName string, headers []string, data [][]string) error {
	recordFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(recordFile)
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	err = writer.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}
