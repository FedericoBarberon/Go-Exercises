package quizgame

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
)

var ErrInvalidCSVFile = fmt.Errorf("invalid QA CSV file")

func GetQAFromFS(fs fs.FS, path string) (QA, error) {
	file, err := fs.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	return getQAFromReader(file)
}

func getQAFromReader(rdr io.Reader) (QA, error) {
	csvReader := csv.NewReader(rdr)
	data, err := csvReader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("failed to reading file: %v", err)
	}

	if !isValidQA(data) {
		return nil, ErrInvalidCSVFile
	}

	return QA(data), nil
}

func isValidQA(data [][]string) bool {
	for _, tuple := range data {
		if len(tuple) != 2 {
			return false
		}
	}

	return true
}
