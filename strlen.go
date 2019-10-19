package student

func StrLen(str string) {

	a := 0
	for _, i := range str {
		if i == i {
			a++
		}
	}
	return a

}
