package words

import (
	"sort"
	"sync"
)

type MostCommonFreqWordsCounter struct {
	num       int
	mu        sync.Mutex // Mutex to protect wordCount
	wordCount map[string]uint
}

func NewMostCommonFreqWordsCounter(num int) *MostCommonFreqWordsCounter {
	return &MostCommonFreqWordsCounter{
		num:       num,
		wordCount: make(map[string]uint),
	}
}

func (c *MostCommonFreqWordsCounter) Add(word WordFreq) {
	c.mu.Lock()
	c.wordCount[word.Word] += word.Freq
	c.mu.Unlock()
}

func (c *MostCommonFreqWordsCounter) Get() []WordFreq {
	var wordFreqs []WordFreq
	c.mu.Lock()
	for word, freq := range c.wordCount {
		wordFreqs = append(wordFreqs, WordFreq{Word: word, Freq: freq})
	}
	c.mu.Unlock()

	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].Freq > wordFreqs[j].Freq
	})

	if c.num >= len(wordFreqs) {
		return wordFreqs
	}

	return wordFreqs[:c.num]
}
