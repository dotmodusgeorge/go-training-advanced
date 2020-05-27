package internal

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
	"bufio"
	"log"
)

func readCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := [][]string{}
	for scanner.Scan() {
		row := scanner.Text()
		content = append(content, strings.Split(row, ","))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return content
}

func mapContent(content [][]string) []map[string]string {
	firstRow := content[0]
	mappedContent := []map[string]string{}
	for _, row := range content[1:] {
		mappedRow := map[string]string{}
		for index, title := range firstRow {
			mappedRow[title] = row[index]
		}
		mappedContent = append(mappedContent, mappedRow)
	}
	return mappedContent
}

func writeJSONFile(content []map[string]string, filePath string, output string) {
	data, err := json.MarshalIndent(content, "", "    ")
	if (err != nil) {
		panic(err)
	}
	filePathParts := strings.Split(filePath, "/")
	fileName := strings.Split(filePathParts[(len(filePathParts) - 1):][0], ".")[0]
	finalPath := fmt.Sprint(output, fileName, ".json")
	err = ioutil.WriteFile(finalPath, data, 0644)

	if (err != nil) {
		panic(err)
	}
}

func HandleCsv(filePath string, output string) {
	content := readCsv(filePath)
	contentMapped := mapContent(content)
	writeJSONFile(contentMapped, filePath, output)
}