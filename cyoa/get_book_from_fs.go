package cyoa

import (
	"encoding/json"
	"errors"
	"io"
	"io/fs"
)

var ErrInvalidJSON = errors.New("invalid json file")

func GetBookFromFS(fs fs.FS, path string) (Book, error) {
	file, err := fs.Open(path)
	if err != nil {
		return nil, errors.New("error opening file")
	}
	defer file.Close()

	return getBookFromReader(file)
}

func getBookFromReader(rdr io.Reader) (Book, error) {
	var book Book
	err := json.NewDecoder(rdr).Decode(&book)
	if err != nil {
		return nil, ErrInvalidJSON
	}

	return book, nil
}
