package handler

import (
	"encoding/csv"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func HandleParsing(dirPath string, queryWords []string) {
	csvFilePaths := readDir(dirPath)

	if len(csvFilePaths) != 0 {
		log.Printf("CSV files to open: %s", csvFilePaths)

		resultRecords := processCsvFiles(csvFilePaths, queryWords)
		writeResult(resultRecords)
	} else {
		log.Printf("Found no CSV files in: %s", dirPath)
	}

}

func readDir(dirName string) []string {
	csvFiles := make([]string, 0)

	err := filepath.Walk(dirName, func(path string, info fs.FileInfo, err error) error {
		if err == nil && isCsvFile(path) {
			csvFiles = append(csvFiles, path)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("FATAL: Error walking through: %s", dirName)
	}

	return csvFiles
}

func isCsvFile(fileName string) bool {
	return filepath.Ext(fileName) == ".csv"
}

func processCsvFiles(fileNames []string, queryWords []string) [][]string {
	resultRecords := make([][]string, 0)

	for _, name := range fileNames {
		file, err := os.Open(name)

		if err != nil {
			log.Fatalf("FATAL: Unable to open file: %s", err.Error())
		}

		fileRecords := readCsvFile(file)
		fileResultRecords := parseFileRecords(fileRecords, queryWords)
		resultRecords = append(resultRecords, fileResultRecords...)
	}

	log.Println("Files were successfully processed")

	return resultRecords
}

func readCsvFile(file *os.File) [][]string {
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatalf("FATAL: Unable to read: %s", err)
	}

	log.Printf("Reading: %s", file.Name())

	return records
}

func parseFileRecords(records [][]string, queryWords []string) [][]string {
	currentResultRecords := make([][]string, 0)

	for _, record := range records {

	out:
		for _, field := range record {
			for _, queryWord := range queryWords {
				if strings.Contains(field, queryWord) {
					currentResultRecords = append(currentResultRecords, record)
					break out
				}
			}

		}
	}

	return currentResultRecords
}

func writeResult(resultRecords [][]string) {
	const resultFileName = "result/output.txt"

	out, _ := os.Create(resultFileName)
	defer out.Close()

	for _, record := range resultRecords {
		out.WriteString(strings.Join(record, ",") + "\n")
	}

	log.Printf("Results were successfully written to: %s\n", resultFileName)
}
