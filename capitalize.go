package student

func Capitalize(s string) string {

	A := []rune(s)

	for i := range s {

		if (A[i] == 32) && (A[i+1] > 96 && A[i+1] < 123) {

			A[i+1] = A[i+1] - rune(32)

		}
	}

	text := string(A)

	return text
}
