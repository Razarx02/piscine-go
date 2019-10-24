package student

func ToUpper(s string) string {

	Size := len(s)

	A := []rune(s)

	for i := 0; i < Size; i++ {

		for x := 97; x <= 1222; x++ {

			if int(A[i]) == x {

				A[i] = rune(x - 32)

			}
		}
	}

	text := string(A)

	return text
}
