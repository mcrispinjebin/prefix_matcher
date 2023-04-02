package main

import (
	"fmt"
	"log"
	"os"
	"prefix_matcher/constants"
	"prefix_matcher/handler"
	"prefix_matcher/models"
	"sync"
)

func init() {
	handler.PreCompute()
}

func main() {
	var sendWg sync.WaitGroup
	var receiveWg sync.WaitGroup
	c := make(chan models.MatcherParam)

	words := []string{"CO", "2Y", "V9", "28", "A6", "EfG", "E"}

	defer func() {
		// To remove the sub files created on pre computation while server stops
		for i := 0; i < constants.BucketsCount; i++ {
			os.Remove(fmt.Sprintf(constants.SubFileNameFormat, i))
		}
	}()

	go func() {
		// Receiving go routine to listen on the channel for output
		receiveWg.Add(len(words))
		for i := 0; i < len(words); i++ {
			result := <-c
			fmt.Printf("Input word is: %s and Matched long word is: %s\n", result.InputWord, result.MatchedWord)
			receiveWg.Done()
		}
	}()

	for i := 0; i < len(words); i++ {
		// i variable is to control both spanning of go routines - max 5 at a time and word picking
		sendWg.Add(1)
		go handler.GetLongestWord(words[i], c, &sendWg)

		//TODO: Read MAX_GOROUTINE_ALLOWED from env rather than constant 5 here
		if i%constants.MaxGoRoutinesAllowed == 0 {
			sendWg.Wait()
		}
	}

	defer func() {
		if r := recover(); r != nil {
			log.Print("panic occurred")
		}
	}()

	receiveWg.Wait()
	close(c)
}
