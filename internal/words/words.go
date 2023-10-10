package words

import (
	"fmt"
	"os"
)

type WordFreq struct {
	Word string
	Freq uint
}

func IsLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func SaveWordFreqsToFile(wordFreqs []WordFreq, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, wf := range wordFreqs {
		line := fmt.Sprintf("%d. %s,%d\n", i+1, wf.Word, wf.Freq)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
