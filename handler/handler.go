package handler

import (
	"prefix_matcher/models"
	"prefix_matcher/usecase"
	"sync"
)

func GetLongestWord(word string, c chan<- models.MatcherParam, wg *sync.WaitGroup) {
	matcher := usecase.NewMatcherUsecase(word)
	matcher.Process(c)
	wg.Done()
}

func PreCompute() {
	preProcess := usecase.InitPreCompute()
	preProcess.PreCompute()
}
