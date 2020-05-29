package internal

import (
	"fmt"
	"os"
	"encoding/json"
	"encoding/csv"
	"io/ioutil"
	"sort"
	"strings"
)

func readJSONFile(filePath string) []map[string]string  {
	jsonFile, err := os.Open(filePath)
	if (err != nil) {
		panic(err)
	}
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var result []map[string]string

	json.Unmarshal(jsonBytes, &result)

	return result
}

func mapRows(content []map[string]string) [][]string {
	keys := make([]string, 0)
	for key := range content[0] {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	rows := make([][]string, 0)
	rows = append(rows, keys)

	for _, obj := range content {
		row := make([]string, 0)
		for _, key := range keys {
			row = append(row, obj[key])
		}
		rows = append(rows, row)
	}
	return rows
}

func writeCsv(content [][]string, filePath string, output string) {
	filePathParts := strings.Split(filePath, "/")
	fileName := strings.Split(filePathParts[(len(filePathParts) - 1):][0], ".")[0]
	finalPath := fmt.Sprint(output, fileName, ".csv")
	f, err := os.OpenFile(finalPath, os.O_WRONLY|os.O_CREATE, 0644)
	if (err != nil) {
		panic(nil)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, thing := range content {
		w.Write(thing)
	}
}

func HandleJSON(filePath string, output string) {
	content := readJSONFile(filePath)
	mappedContent := mapRows(content)
	writeCsv(mappedContent, filePath, output)
}





