package student

func FirstRune(s string) rune {

	W := 0

	for i := -50; i <= 300; i++ {

		if int(s[0]) == i {

			W = i

		}

	}

	return rune(W)
}
