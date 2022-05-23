package chapter1

import "strings"

func URLify(input string, length int) string {
	count := 0
	for _, v := range input {
		count++
		if v == ' ' {
			// Extra "20"
			count += 2
		}
	}

	// How about `make([]rune, length, extra)` ?
	builder := make([]rune, count)
	i := 0
	for _, v := range input {
		if v == ' ' {
			builder[i] = '%'
			builder[i+1] = '2'
			builder[i+2] = '0'
			i += 3
		} else {
			builder[i] = v
			i++
		}
	}

	return string(builder)
}

func URLifyWithBuilder(input string, length int) string {
	sb := strings.Builder{}
	for _, v := range input {
		if v == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteRune(v)
		}
	}
	return sb.String()
}
