package chapter1

import (
	"strings"
)

func IsPermutatedPalindrome(input string) bool {
	freq := map[rune]int{}
	for _, v := range strings.ToLower(input) {
		if v == ' ' {
			// Ignore space
			continue
		}

		if _, ok := freq[v]; ok {
			freq[v]++
		} else {
			freq[v] = 1
		}
	}

	oddFlag := false
	for _, v := range freq {
		if v%2 != 0 {
			if oddFlag {
				return false
			}
			oddFlag = true
		}
	}
	return true
}
