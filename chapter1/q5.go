package chapter1

func isReplacedOnce(val1, val2 string) bool {
	edited := false
	val2Rune := []rune(val2)
	for i, v := range val1 {
		if val2Rune[i] != v {
			if edited {
				return false
			}
			edited = true
		}
	}
	return true
}

func isInsertOnce(val1, val2 string) bool {
	// Check insert once, so implicitly assumed len(val1) < len(val2)
	cursor := 0

	val2Rune := []rune(val2)

	edited := false
	for _, v := range val1 {
		if v != val2Rune[cursor] {
			if edited {
				return false
			}
			cursor += 2
		} else {
			cursor++
		}
	}
	return true
}

func IsOneStep(val1, val2 string) bool {
	if len(val1) == len(val2) {
		// replace
		return isReplacedOnce(val1, val2)
	}

	if len(val2)-len(val1) == 1 {
		// Insertion
		return isInsertOnce(val1, val2)
	}

	if len(val1)-len(val2) == 1 {
		// Deletion
		return isInsertOnce(val2, val1)
	}

	return false
}
