package chapter1

import "strings"

func IsPermutation(val1, val2 string) bool {
	target := val2
	for _, v := range val1 {
		idx := strings.IndexRune(target, v)
		if idx == -1 {
			return false
		}
		target = target[:idx] + target[idx+1:]
	}
	if len(target) != 0 {
		return false
	}
	return true
}
