package usecase

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"prefix_matcher/constants"
	"sort"
	"sync"
)

type PreComputation interface {
	SortAndStorePrefixWords(categorisedMap map[int]*[]string, sortKey int, wg *sync.WaitGroup)
	SubCategorizePrefixes(prefixWord string, subPrefixesMap map[int]*[]string)
	PreCompute()
}

type preComputer struct{}

func InitPreCompute() PreComputation {
	return preComputer{}
}

func (p preComputer) PreCompute() {
	var (
		wg sync.WaitGroup
	)
	categorisedPrefixMap := make(map[int]*[]string)

	file, err := os.Open(constants.SourceFilePath)
	if err != nil {
		log.Fatalf("error occurred in opening the mentioned file -  %+v", err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prefixWord := scanner.Text()
		p.SubCategorizePrefixes(prefixWord, categorisedPrefixMap)
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("error in scanning the files - %+v", err.Error())
		return
	}

	for i := 0; i < constants.BucketsCount; i++ {
		wg.Add(1)
		go p.SortAndStorePrefixWords(categorisedPrefixMap, i, &wg)
	}
	wg.Wait()

}

func (p preComputer) SortAndStorePrefixWords(categorisedMap map[int]*[]string, sortKey int, wg *sync.WaitGroup) {
	defer wg.Done()
	sortItems := *categorisedMap[sortKey]
	sort.Slice(sortItems, func(i, j int) bool {
		return sortItems[i] > sortItems[j]
	})

	fileName := fmt.Sprintf(constants.SubFileNameFormat, sortKey)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("error occurred in writing to a file")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, val := range sortItems {
		_, err = writer.WriteString(fmt.Sprintf("%s\n", val))
		if err != nil {
			log.Fatalf("error in writing a prefix word to sub Categorised file - %+v", val)
		}
	}
	writer.Flush()
}

func (p preComputer) SubCategorizePrefixes(prefixWord string, subPrefixesMap map[int]*[]string) {
	prefixBucketPage := FindBucket(prefixWord)
	if prefixBucketPage >= 0 && prefixBucketPage <= constants.BucketsCount {
		if subPrefixesMap[prefixBucketPage] == nil || len(*subPrefixesMap[prefixBucketPage]) == 0 {
			subPrefixesMap[prefixBucketPage] = &[]string{prefixWord}
		} else {
			*subPrefixesMap[prefixBucketPage] = append(*subPrefixesMap[prefixBucketPage], prefixWord)
		}
	} else {
		log.Fatalf("bucket Page calculated is not a valid value - %+v for prefix - %+v", prefixBucketPage, prefixWord)
		return
	}
	return
}
