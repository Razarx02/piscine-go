package student

func NRune(s string, n int) rune {

	A := []rune(s)

	for index, word := range A {

		if A[n-1] == word {
			return index
		}
	}

	return 0
}
