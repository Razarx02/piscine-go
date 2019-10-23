package student

func LastRune(s string) rune {

	A := []rune(s)

	for index, word := range A {

		if index == len(A-1) {
			return word
		}
	}

	return 0
}
