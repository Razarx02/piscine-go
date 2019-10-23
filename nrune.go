package student

func NRune(s string, n int) rune {

	A := []rune(s)

	for _, word := range A {

		if A[n-1] == word {
			return word
		}
	}

	return 0
}
