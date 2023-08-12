package handler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func HandleParsing(dirPath string, queryWords []string) {
	files := getFilesFromDirectory(dirPath)

	csvFileNames := getCsvFileNames(files)

	log.Printf("CSV files to open: %s\n", csvFileNames)

	resultRecords := processCsvFiles(csvFileNames, dirPath, queryWords)

	writeResult(resultRecords)
}

func getFilesFromDirectory(dirPath string) []os.FileInfo {

	dir, err := os.Open(dirPath)
	if err != nil {
		log.Fatal("Unable to read directory "+dirPath, err)
	}

	log.Printf("Successfully opened dir %s", dirPath)

	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal("Unable to read files in directory "+dirPath, err)
	}

	return files
}

func getCsvFileNames(files []os.FileInfo) []string {
	csvFiles := make([]string, 0)

	for _, file := range files {
		log.Printf("Current file: %s", file.Name())
		if isCsvFile(file.Name()) {
			csvFiles = append(csvFiles, file.Name())
		}
	}

	return csvFiles
}

func isCsvFile(fileName string) bool {
	return filepath.Ext(fileName) == ".csv"
}

func processCsvFiles(fileNames []string, dirPath string, queryWords []string) [][]string {
	resultRecords := make([][]string, 0)

	for _, name := range fileNames {
		file, err := os.Open(fmt.Sprintf("%s/%s", dirPath, name))
		if err != nil {
			log.Fatal("Unable to open file "+name, err)
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
		log.Fatal("Unable to open file "+file.Name(), err)
	}

	log.Printf("Reading %s", file.Name())

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
	out, _ := os.Create("result/output.txt")

	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal("Unable to create result file")
		}
	}(out)

	for _, record := range resultRecords {
		out.WriteString(strings.Join(record, ",") + "\n")
	}
}
