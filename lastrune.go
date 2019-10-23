package student

func LasRune(s string, n int) rune {

	A := []rune(s)

	for index, word := range A {

		if index == len(A-1) {
			return word
		}
	}

	return 0
}
