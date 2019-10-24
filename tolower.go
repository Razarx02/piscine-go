package student

func ToLower(s string) string {

	Size := len(s)

	A := []rune(s)

	for i := 0; i < Size; i++ {

		for x := 65; x <= 90; x++ {

			if int(A[i]) == x {

				A[i] = rune(x + 32)

			}
		}
	}

	text := string(A)

	return text
}
