package student

func MakeRange(min int, max int) []int {

	var forz []int = make([]int, 0)

	if min == max || min > max {
		return forz
	}
	sred := max - min
	var Buffer []int = make([]int, sred)
	index := 0

	for i := min; i < max; i++ {

		Buffer[index] = i
		index++

	}

	return Buffer

}
