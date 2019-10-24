package student

func MakeRange(min int, max int) []int {

	if min == max || min > max {
		return []int(nil)
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
