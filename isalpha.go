package student

func IsAlpha(str string) bool {

	xRune := []rune(str)
	Cat := 0

	for range str {
		Cat++
	}

	A := true

	for i := 0; i < Cat; i++ {

		if xRune[i] == 32 {

			A = false

		}
	}

	return A
}
