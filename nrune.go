package student

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {

	S := []rune(s)

	return S[n-1]

}
