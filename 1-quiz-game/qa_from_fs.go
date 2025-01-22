package quizgame

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
)

var ErrInvalidCSVFile = fmt.Errorf("invalid QA CSV file")

func GetQAPairsFromFS(fs fs.FS, path string) ([]QAPair, error) {
	file, err := fs.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	qaPairs, err := getQAPairsFromReader(file)

	if err == ErrInvalidFormat {
		return nil, ErrInvalidCSVFile
	}

	return qaPairs, nil
}

func getQAPairsFromReader(rdr io.Reader) ([]QAPair, error) {
	csvReader := csv.NewReader(rdr)
	data, err := csvReader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("failed to reading file: %v", err)
	}

	return CastQaPairs(data)
}
