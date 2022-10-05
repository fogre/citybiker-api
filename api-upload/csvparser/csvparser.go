package csvparser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

/*
Opens file in inputFileSrc for reading
and creates an outputFile to outPutFileScr for writing
*/
func initFiles(inputFileSrc string, outputFileSrc string) (inputFile *os.File, outputFile *os.File) {
	inputFile, inputError := os.Open(inputFileSrc)
	if inputError != nil {
		log.Fatalln(inputError)
	}

	outputFile, outputError := os.Create(outputFileSrc)
	if outputError != nil {
		log.Fatalln(outputError)
	}

	return inputFile, outputFile
}

// Returns csv.Writer and writes database Model type headers to the outputFile
func initCsvWriter(writeFile *os.File, headers []string) *csv.Writer {
	writer := csv.NewWriter(writeFile)
	if err := writer.Write(headers); err != nil {
		log.Fatalln(err)
	}
	return writer
}

// Type for func validating a csv record before writing it to outputFile
type csvValidRecordWriter func([]string, *csv.Writer)

// Writes a valid Station model to csv output file
func writeValidStation(record []string, writer *csv.Writer) {
	if record[1] != "" &&
		record[2] != "" &&
		record[5] != "" &&
		record[11] != "" &&
		record[12] != "" {

		writer.Write([]string{
			record[1],
			record[2],
			record[5],
			record[11],
			record[12],
		})
	}
}

// Parses through the input CSV file, and writes only valid fields to output CSV
func parseAndWriteOutputCsv(reader *csv.Reader, writer *csv.Writer, validWriter csvValidRecordWriter) {
	//Read first row of the input (headers) so it's not written to output
	if _, err := reader.Read(); err != nil {
		log.Fatalln(err)
	}
	//parse csv records
	for {
		record, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		validWriter(record, writer)
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatalln(err)
	}
}

func ParseStations(inputFileSrc string, outputFileSrc string) {
	inputFile, outputFile := initFiles(inputFileSrc, outputFileSrc)
	defer inputFile.Close()
	defer outputFile.Close()

	reader := csv.NewReader(inputFile)

	headers := []string{"Id", "Name", "Address", "CoordinateX", "CoordinateY"}
	writer := initCsvWriter(outputFile, headers)

	parseAndWriteOutputCsv(reader, writer, writeValidStation)
	fmt.Printf("Validated Station output csv created: %s\n", outputFileSrc)
}

/*
func ParseTrips(inputFileSrc string, outputFileSrc string) {
	inputFile, outputFile := initFiles(inputFileSrc, outputFileSrc)
	defer inputFile.Close()
	defer outputFile.Close()

	reader := csv.NewReader(inputFile)

	headers := []string{"Id", "Name", "Address", "CoordinateX", "CoordinateY"}
	writer := initCsvWriter(outputFile, headers)

	parseAndWriteOutputCsv(reader, writer, writeValidStation)
	fmt.Printf("Validated Station output csv created: %s\n", outputFileSrc)
}*/
