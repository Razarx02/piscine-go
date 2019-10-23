package student

func NRune(s string, n int) rune {

	A := []rune(s)

	for _, word := range A {

		if A[n-1] == word {
			return int(A[n-1])
		}
	}

	return 0
}
