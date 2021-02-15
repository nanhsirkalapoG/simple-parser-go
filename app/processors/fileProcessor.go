package processors

import (
	"fmt"
	"parser/app/readers"
	"parser/app/writers"
	"path"
)

var OutputDirectory = "resources"

type FileProcessor struct {
	
}


// Read, parse and generate a standard format data. Write the data into CSV files
func (fp FileProcessor) ProcessFile(fileName string) error {
	fmt.Println("Processing File: " + fileName)
	// Get a CSV reader to read the data
	reader, err := readers.ReadFile(fileName)
	if err != nil {
		return err
	}

	// Process the CSV file records and create a standard format
	result, err := ProcessRecords(reader.Records, reader.Headers, reader.HeaderMap)
	if err != nil {
		return err
	}

	//Write the processed data into CSV files
	err = fp.WriteOutputToFile(result, fileName)
	if err != nil {
		return err
	}
	fmt.Println("Processed Records: ", len(result.ValidRecords))
	fmt.Println("Unprocessed Records: ", len(result.InValidRecords))
	fmt.Println("---------------------------------------")

	return nil
}

// Write data to files
func (fp FileProcessor) WriteOutputToFile(result Result, fileName string) error {
	_, file := path.Split(fileName)
	err := writers.WriteFile(fp.GetValidFileName(file), result.ValidRecordsHeader, result.ValidRecords)
	if err != nil {
		return err
	}

	err = writers.WriteFile(fp.GetInValidFileName(file), result.InValidRecordsHeader, result.InValidRecords)
	if err != nil {
		return err
	}

	return nil
}

func (fp FileProcessor) GetValidFileName(fileName string) string {
	return fmt.Sprintf("%s/valid_%s", OutputDirectory, fileName)
}

func (fp FileProcessor) GetInValidFileName(fileName string) string {
	return fmt.Sprintf("%s/invalid_%s", OutputDirectory, fileName)
}
