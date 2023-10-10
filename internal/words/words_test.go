package words

import (
	"testing"
)

func TestIsLetter(t *testing.T) {
	for r := 'a'; r <= 'z'; r++ {
		if !IsLetter(r) {
			t.Errorf("Expected %c to be a letter, but it wasn't", r)
		}
	}

	for r := 'A'; r <= 'Z'; r++ {
		if !IsLetter(r) {
			t.Errorf("Expected %c to be a letter, but it wasn't", r)
		}
	}

	nonLetters := []rune{'1', '2', '@', ' ', '!', '$', '&'}
	for _, r := range nonLetters {
		if IsLetter(r) {
			t.Errorf("Expected %c to not be a letter, but it was", r)
		}
	}
}
