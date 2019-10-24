package student

func IsAlpha(str string) bool {

	xRune := []rune(str)
	Cat := 0

	for range str {
		Cat++
	}

	A := true

	for i := 0; i < Cat; i++ {

		if xRune[i] > 48 || (xRune[i] > 39 && xRune[i] < 41) || (xRune[i] > 90 && xRune[i] < 61) || (xRune[i] > 122) {

			A = false

		}
	}

	return A
}
