package usecase

import (
	"testing"
)

func TestWordMatcher_AppropriateByBinarySearch(t *testing.T) {
	tests := []struct {
		name      string
		inputWord string
		wordList  []string
		expResult int
	}{
		{"last_element_search", "abcd", []string{"decd", "cdce", "bcde", "abcd"}, 3},
		{"unhappy_sort_different_order", "abcd", []string{"abcd", "bcde", "cdce", "decd"}, -1},
		{"unhappy_not_in_list", "abcd", []string{"cbcd", "dcde", "edce", "fecd"}, -1},
		{"happy_all_are_same_char", "abcd", []string{"abcd", "acde", "adce", "aecd"}, 2},
		{"first_element_search", "decd", []string{"decd", "cdce", "bcde", "abcd"}, 0},
	}

	for _, subtest := range tests {
		t.Run(subtest.name, func(t *testing.T) {
			w := NewMatcherUsecase(subtest.inputWord)
			result := w.AppropriateByBinarySearch(subtest.wordList)
			if result != subtest.expResult {
				t.Logf("AppropriateByBinarySearch() : FAILED, expected line num %v but got result %v", subtest.expResult, result)
			} else {
				t.Logf("AppropriateByBinarySearch() : PASSED, expected line num %v and got result %v", subtest.expResult, result)
			}
		})
	}
}

func TestWordMatcher_FindWordByProbing(t *testing.T) {
	subtests := []struct {
		name      string
		word      string
		wordsList []string
		i         int
		expResult string
	}{
		{"find_word_at_first", "de", []string{"eefnkdsbfsn", "defgh", "deb", "dec", "dea"}, 4, "defgh"},
		{"no_match", "de", []string{"eefnkdsbfsn", "ddefgh", "ddeb", "ddec", "ddea"}, 4, "-1"},
		{"decrement_search", "de", []string{"eefnkdsbfsn", "defgh", "ddb", "dcc", "dca"}, 4, "defgh"},
		{"incremental_search", "dc", []string{"eefnkdsbfsn", "defgh", "ddb", "dcc", "dca"}, 1, "dcc"},
	}

	for _, subtest := range subtests {
		w := NewMatcherUsecase(subtest.word)
		t.Run(subtest.name, func(t *testing.T) {
			w.FindWordByProbing(subtest.wordsList, subtest.i)
		})
	}
}
