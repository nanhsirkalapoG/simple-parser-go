package main

import (
	"fmt"
	"os"
	"parser/app/processors"
)

func main() {
	fmt.Println("Starting File Parsing")
	// Read arguments
	fileList := os.Args[1:]
	fmt.Println(fmt.Sprintf("Total number of files: %d", len(fileList)))
	fileProcessor := processors.FileProcessor{}

	// Process each file
	for _, file := range fileList {
		if err := fileProcessor.ProcessFile(file); err != nil {
			fmt.Println("could not process file: " + file, err)
		}
	}
	fmt.Println("Completed File Parsing")
}
