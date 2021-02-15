package readers

import (
	"encoding/csv"
	"io"
	"os"
)

type Reader struct {
	Headers []string
	Records [][]string
	HeaderMap map[string]int
}

// Open the file and create a reader to read the data
func ReadFile(fileName string) (Reader, error) {
	reader := Reader{}
	var records [][]string

	recordFile, err := os.Open(fileName)
	if err != nil {
		return reader, err
	}

	csvReader := csv.NewReader(recordFile)
	headers, err := csvReader.Read()
	if err != nil {
		return reader, err
	}
	reader.Headers = headers
	//fmt.Printf("Headers : %v \n", headers)

	headerMap := make(map[string]int)
	for i, f := range headers {
		headerMap[f] = i + 1
	}
	reader.HeaderMap = headerMap

	records, err = readRecords(csvReader)
	reader.Records = records

	return reader, err
}

// Read records one by one
func readRecords(reader *csv.Reader) ([][]string, error) {
	var records [][]string

	for i := 0; ; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			return records, err
		}
		records = append(records, record)
		//fmt.Printf("Row %d : %v \n", i, record)
	}

	return records, nil
}
