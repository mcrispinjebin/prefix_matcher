package usecase

import "testing"

func TestFindBucket(t *testing.T) {
	testCases := []struct {
		name            string
		inputParam      string
		outputBucketNum int
	}{
		{"first_bucket_lower_case", "abcd", 0},
		{"first_bucket_upper_case", "Abcd", 0},
		{"digit_last_bucket", "7jhsjdek", 5},
		{"special_char_last_bucket", "!sdfn", 5},
	}

	for _, subtest := range testCases {
		t.Run(subtest.name, func(t *testing.T) {
			result := FindBucket(subtest.inputParam)
			if result != subtest.outputBucketNum {
				t.Logf("FindBucket() : FAILED, expected  bucket num %v but got result %v", subtest.outputBucketNum, result)
			} else {
				t.Logf("FindBucket() : PASSED, expected  bucket num %v and got result %v", subtest.outputBucketNum, result)
			}
		})
	}
}
