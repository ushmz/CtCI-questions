package chapter1

import (
	"strconv"
)

func CompressText(input string) string {
	buf := input[:1]
	count := 1
	var comp string
	for _, v := range input[1:] {
		if string(v) == buf {
			count++
			continue
		}
		comp += (string(buf) + strconv.Itoa(count))
		if len(comp) >= len(input) {
			return input
		}
		buf = string(v)
		count = 1
	}
	return comp + (string(buf) + strconv.Itoa(count))
}
