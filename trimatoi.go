package student

func TrimAtoi(s string) int {

	Nomer := 0
	Chekminus := 1

	for _, i := range s {

		if i >= '0' && i <= '9' {

			N := int(i - 48)

			Nomer = Nomer*10 + N

		} else if i == '-' && Nomer == 0 {

			Chekminus = -1

		}

	}

	return Nomer * Chekminus
}
