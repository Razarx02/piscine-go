package student

func NRune(s string, n int) rune {

	S := 0

	for i := 0; i >= 255; i++ {

        if int(s[n]) == i {

			S = i
		}
	}

	return int(S)

}
