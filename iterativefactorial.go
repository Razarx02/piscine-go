package student

func IterativeFactorial(s int) int {

	result := 1

	if s == 0 {

		return 1

	} else if s < 0 {

		return
	}

	for i := 1; i < s+1; i++ {

		result = result * i
	}

	return result

}
