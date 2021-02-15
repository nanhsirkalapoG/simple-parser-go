# File Parser

## Description
    Parse a given file and create a standard csv file
    - Input: single CSV file(i.e input.csv)
    - Ouput: two CSV file(i.e valid.csv, invalid.csv)
    
## How to run the application?
###Build the application
    `go build application.go`
###Execute the binary
    `./application file1 [file2 file3 ...`
    Give multiple files separate by space
###Output
    Command line output
        - Input file name followed by
            processed and unprocessed records count
    Files
        - Two files will be created for each input file
            valid_filename.csv and invalid_filename.csv
        
## Sample output
        
###Sample command
        ./application resources/sample1.csv resources/sample2.csv
        
###Sample console output
        Starting File Parsing
        Total number of files: 2
        Processing File: resources/sample1.csv
        error processing record: [Alfred Donald  $11.5 4], error: invalid email address
        Processed Records:  4
        Unprocessed Records:  1
        ---------------------------------------
        Processing File: resources/sample2.csv
        Processed Records:  5
        Unprocessed Records:  0
        ---------------------------------------
        Completed File Parsing
