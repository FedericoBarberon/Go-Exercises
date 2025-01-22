package quizgame

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

var ErrInvalidCSVFile = errors.New("provided an invalid QA CSV file")
var ErrNotCSVFile = errors.New("provided a non csv file")

func GetQAPairsFromFS(fs fs.FS, path string) ([]QAPair, error) {
	if !strings.HasSuffix(path, ".csv") {
		return nil, ErrNotCSVFile
	}

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
