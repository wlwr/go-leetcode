package longest_substring_without_repeating_characters

// Given a string, find the length of the longest substring without repeating characters.

// Examples:

// Given "abcabcbb", the answer is "abc", which the length is 3.

// Given "bbbbb", the answer is "b", with the length of 1.

// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	longest := 0
	indexMap := make(map[string]int)
	start := 0
	for i, r := range s {
		char := string(r)
		if index, ok := indexMap[char]; ok && index >= start {
			start = index + 1
		}
		indexMap[char] = i
		tmpLen := i - start + 1
		if tmpLen > longest {
			longest = tmpLen
		}

	}
	return longest
}

// func lengthOfLongestSubstring(s string) int {
// 	longest := 0
// 	indexMap := make(map[string]int)
// 	tmpLen := 0
// 	for i, r := range s {
// 		char := string(r)
// 		if _, ok := indexMap[char]; ok {
// 			if tmpLen > longest {
// 				longest = tmpLen
// 			}
// 			indexMap = make(map[string]int)
// 			tmpLen = 0
// 		}
// 		tmpLen++
// 		indexMap[char] = i
// 		if tmpLen > longest {
// 			longest = tmpLen
// 		}
// 	}
// 	return longest
// }
