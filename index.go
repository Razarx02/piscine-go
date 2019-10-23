package student

func Index(s string, toFind string) int {

	a := 0
	c := 0

	for range toFind {
		a++
	}

	if a != 0 {
		buffer := []rune(toFind)
		for i, w := range s {
			if (i == i) && (buffer[0] == w && buffer[0] != 10) {
				c++
				return i
			}
		}
		return -1
	}

	return 0
}
