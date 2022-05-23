package chapter1

func IsAllUnique(sentence string) bool {
	setMap := map[rune]int{}
	for _, v := range sentence {
		count, ok := setMap[v]
		if ok {
			return false
		}
		setMap[v] = (count + 1)
	}
	return true
}
