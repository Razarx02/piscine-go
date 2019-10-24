package student

func AppendRange(min int, max int) []int {

	var Buffer []int

	if min == max || min > max {
		return Buffer
	}
	for i := min; i < max; i++ {

		Buffer = append(Buffer, i)
	}

	return Buffer

}
