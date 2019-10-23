package student

func FirstRune(s string) rune {

	W := 0

	for i := 65; i <= 90; i++ {

		if int(s[0]) == i {

			W = i

		}

	}

	return rune(W)
}
