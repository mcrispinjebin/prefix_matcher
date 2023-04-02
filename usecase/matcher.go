package usecase

import (
	"fmt"
	"io"
	"log"
	"os"
	"prefix_matcher/constants"
	"prefix_matcher/models"
	"strings"
)

type wordMatcher struct {
	MatchParam models.MatcherParam
}

func NewMatcherUsecase(inputWord string) WordMatcher {
	return &wordMatcher{MatchParam: models.MatcherParam{
		InputWord:   inputWord,
		MatchedWord: "-1",
	}}
}

type WordMatcher interface {
	Process(c chan<- models.MatcherParam)
	ReadAppropriateSubFile() []string
	FindExactLongWord(wordsList []string, startValue, incrementalValue int)
	FindWordByProbing(wordsList []string, i int)
	AppropriateByBinarySearch(wordsList []string) int
}

func (w *wordMatcher) Process(c chan<- models.MatcherParam) {
	wordsList := w.ReadAppropriateSubFile()
	defer func() {
		c <- w.MatchParam
	}()

	// find sea of Words by matching only the initial character with log(n) time complexity
	lineNum := w.AppropriateByBinarySearch(wordsList)
	if lineNum == -1 {
		return
	}

	// find exact longest word by linear search after matching first character - three cases here
	w.FindWordByProbing(wordsList, lineNum)
}

func (w *wordMatcher) ReadAppropriateSubFile() []string {
	s := w.MatchParam.InputWord
	bucketPage := FindBucket(s)
	fileName := fmt.Sprintf(constants.SubFileNameFormat, bucketPage)
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("[ERROR] error in opening the sub processed file with name - %+v", fileName)
	}
	defer file.Close()
	rawBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("[ERROR] error in reading the sub processed file with name - %+v", fileName)
	}

	wordsList := strings.Split(string(rawBytes), "\n")
	return wordsList
}

func (w *wordMatcher) FindExactLongWord(wordsList []string, startValue, incrementalValue int) {
	toMatchWord := w.MatchParam.InputWord
	maxWord := w.MatchParam.MatchedWord
	for startValue >= 0 && startValue < len(wordsList) && strings.HasPrefix(wordsList[startValue], toMatchWord) {
		if len(wordsList[startValue]) > len(maxWord) {
			w.MatchParam.MatchedWord = wordsList[startValue]
		}
		startValue = startValue + incrementalValue
	}
}

func (w *wordMatcher) FindWordByProbing(wordsList []string, i int) {
	word := w.MatchParam.InputWord
	if strings.HasPrefix(wordsList[i], word) {
		// do search up and down
		w.MatchParam.MatchedWord = wordsList[i]
		j := i - 1
		w.FindExactLongWord(wordsList, j, -1)

		j = i + 1
		w.FindExactLongWord(wordsList, j, 1)
		return
	}

	// matched line is lower since its ascii is less
	if word > wordsList[i][:len(word)] {
		j := i - 1
		for word > wordsList[j][:len(word)] && !strings.HasPrefix(wordsList[j], word) {
			j -= 1
		}
		if strings.HasPrefix(wordsList[j], word) {
			w.MatchParam.MatchedWord = wordsList[j]
			w.FindExactLongWord(wordsList, j, -1)
		}

	} else {
		// matched line is upper since its ascii is high
		j := i + 1
		for word < wordsList[j][:len(word)] && !strings.HasPrefix(wordsList[j], word) {
			j += 1
		}
		if strings.HasPrefix(wordsList[j], word) {
			w.MatchParam.MatchedWord = wordsList[j]
			w.FindExactLongWord(wordsList, j, 1)
		}
	}
}

func (w *wordMatcher) AppropriateByBinarySearch(wordsList []string) int {
	initialAscii := w.MatchParam.InputWord[0]
	i := len(wordsList) / 2
	for i >= 0 {
		if wordsList[i][0] == initialAscii {
			break
		}

		if i <= 0 || i >= len(wordsList)-1 {
			return -1
		}

		if wordsList[i][0] > initialAscii {
			i += (len(wordsList) - i) / 2
		} else {
			i -= (len(wordsList) - i) / 2
		}
	}
	return i
}
