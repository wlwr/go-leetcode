package longest_substring_without_repeating_characters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []string{"abcabcbb", "bbbbb", "pwwkew", "abcdef", "abba", "dvdf"}
	results := []int{3, 1, 3, 6, 2, 3}
	for index, test := range tests {
		if res := lengthOfLongestSubstring(test); results[index] != res {
			t.Fatalf("%d - %d\n", res, index)
		}
	}
}
