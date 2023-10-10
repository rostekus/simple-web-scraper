package words

import (
	"testing"
)

func TestMostCommonFreqWordsCounter(t *testing.T) {
	counter := NewMostCommonFreqWordsCounter(3)

	words := []WordFreq{
		{"Python", 5},
		{"C++", 3},
		{"Python", 2},
		{"Java", 4},
		{"Go", 10},
	}

	for _, word := range words {
		counter.Add(word)
	}

	mostCommon := counter.Get()

	expected := []WordFreq{
		{"Go", 10},
		{"Python", 7},
		{"Java", 4},
	}

	if len(mostCommon) != len(expected) {
		t.Errorf("Expected %d most common words, but got %d", len(expected), len(mostCommon))
	}

	for i, word := range mostCommon {
		if word != expected[i] {
			t.Errorf("Expected word %s with freq %d at position %d, but got %s with freq %d",
				expected[i].Word, expected[i].Freq, i, word.Word, word.Freq)
		}
	}
}
