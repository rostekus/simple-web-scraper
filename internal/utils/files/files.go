package files

import (
	"bufio"
	"os"
)

type UrlFileReader struct {
	Filename string
}

func New(filename string) *UrlFileReader {
	return &UrlFileReader{Filename: filename}
}

func (r *UrlFileReader) Iterator() (*UrlIterator, error) {
	file, err := os.Open(r.Filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	return &UrlIterator{file: file, scanner: scanner}, nil
}

type UrlIterator struct {
	file    *os.File
	scanner *bufio.Scanner
}

func (it *UrlIterator) Next() (string, bool) {
	if it.scanner.Scan() {
		return it.scanner.Text(), true
	}
	it.file.Close()
	return "", false
}
