package files

import (
	"bufio"
	"fmt"
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

func (it *UrlIterator) Next() (string, error) {
	if it.scanner.Scan() {
		return it.scanner.Text(), nil
	}
	it.file.Close()
	return "", fmt.Errorf("the end of file")

}
