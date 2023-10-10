package scraper

import (
	"io"
	"log/slog"
	"net/http"
	"strings"

	w "github.com/rostekus/simple-web-scraper/internal/words"
	"golang.org/x/net/html"
)

// define interface where it will be used
type UrlIterator interface {
	Next() (string, bool)
}

type Scraper struct {
	results chan<- w.WordFreq
}

func New(logger *slog.Logger, results chan<- w.WordFreq) *Scraper {
	return &Scraper{
		results: results,
	}
}

// function calculates frequency of occurence of the words in given url and sends it the channel
func (s *Scraper) CalcFreqWords(url string) (map[string]uint, error) {

	responseBody, err := s.getBody(url)
	if err != nil {
		return nil, err
	}
	defer responseBody.Close()

	tokenizer := html.NewTokenizer(responseBody)
	var text string

forLoop:

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			break forLoop
		case html.TextToken:
			text += tokenizer.Token().Data
		}
	}

	words := strings.FieldsFunc(text, func(r rune) bool {
		return !w.IsLetter(r)
	})

	// TODO estimate map size for optimization
	wordFreq := make(map[string]uint)

	for _, word := range words {
		word = strings.ToLower(strings.Trim(word, ".,!?()[]{}\"'"))
		wordFreq[word]++

	}
	return wordFreq, nil
}

func (s *Scraper) getBody(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, err
	}
	return response.Body, nil

}
