package main

import (
	"csv_parser/handler"
)

func main() {
	dirPath := "files"
	queryWords := []string{"test", "world"}

	handler.HandleParsing(dirPath, queryWords)
}
