package usecase

import (
	"prefix_matcher/constants"
	"strings"
)

func FindBucket(prefixWord string) (prefixBucketPage int) {
	prefixBucketPage = int(strings.ToLower(prefixWord)[0]-97) / (constants.BucketsCount - 1)
	if int(strings.ToLower(prefixWord)[0]) < 97 {
		prefixBucketPage = constants.BucketsCount - 1
	}
	return
}
