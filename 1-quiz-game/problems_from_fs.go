package quizgame

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

var ErrInvalidCSVFile = errors.New("provided an invalid problems CSV file")
var ErrNotCSVFile = errors.New("provided a non csv file")

func GetProblemsFromFS(fs fs.FS, path string) ([]Problem, error) {
	if !strings.HasSuffix(path, ".csv") {
		return nil, ErrNotCSVFile
	}

	file, err := fs.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	problems, err := getProblemsFromReader(file)

	if err == ErrInvalidFormat {
		return nil, ErrInvalidCSVFile
	}

	return problems, nil
}

func getProblemsFromReader(rdr io.Reader) ([]Problem, error) {
	csvReader := csv.NewReader(rdr)
	data, err := csvReader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("failed to reading file: %v", err)
	}

	return CastProblems(data)
}
