package student

func Fibonacci(Number int) int {

	toHelp := 0

	if Number < 0 {

		return -1

	}

	if Number < 0 || Number > 12 {

		return 0
	}

	for i := 0; i <= Number; i++ {

		X1 := i - 1
		X2 := i - 2

		if X1+X2 == i {

			toHelp = i

		} else {

			X1 = 0
			X2 = 0

		}

	}

	return toHelp
}
