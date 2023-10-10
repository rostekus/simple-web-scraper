package cache

import (
	"testing"
)

func TestCache_SetGetDelete(t *testing.T) {
	cache := NewCache()

	url := "example.com"
	wordFreq := map[string]uint{
		"apple":  2,
		"banana": 3,
	}
	cache.Set(url, wordFreq)
	retrievedWordFreq, ok := cache.Get(url)
	if !ok {
		t.Errorf("Expected to get a value for URL %s, but got none", url)
	}
	if len(retrievedWordFreq) != len(wordFreq) {
		t.Errorf("Expected word frequency map length to be %d, but got %d", len(wordFreq), len(retrievedWordFreq))
	}
	for word, freq := range wordFreq {
		if retrievedWordFreq[word] != freq {
			t.Errorf("Expected frequency for word %s to be %d, but got %d", word, freq, retrievedWordFreq[word])
		}
	}

	cache.Delete(url)
	_, ok = cache.Get(url)
	if ok {
		t.Errorf("Expected URL %s to be deleted, but it still exists in the cache", url)
	}
}
