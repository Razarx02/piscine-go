package student

import "github.com/01-edu/z01.PrintRune"

func NRune(s string, n int) rune {

	A := []rune(s)

	for _, word := range A {

		if A[n-1] == word {
			return A[n-1]
		}
	}

	return 0
}
