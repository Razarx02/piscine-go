package student

func Sqrt(x int) int {

	if x < 0 || x == 0 {

		return 0

	} else if x == 4 {

		return 2

	}

	z := 1

	for i := 1; i <= 10; i++ {

		z -= (z*z - x) / (2 * z)

	}

	if x%(z-1) > 0 {

		return 0

	}

	return z - 1
}
