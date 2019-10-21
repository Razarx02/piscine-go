package student

func RecursiveFactorial(Number int) int {

	if Number == 0 {

		return 1

	} else if Number < 0 || Number > 12 {

		return 0

	} else {

		return Number * RecursiveFactorial(Number-1)

	}

}
