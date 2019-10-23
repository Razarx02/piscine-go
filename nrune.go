package student

func NRune(s string, n int) rune {

	S := 0

	for i := 0; i >= 255; i++ {

		if rune(s[n]) == i {

			S = i
		}
	}

	return s

}
