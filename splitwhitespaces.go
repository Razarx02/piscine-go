package student

func SplitWhiteSpaces(str string) []string {

	lol := 0

	for range str {
		lol++
	}

	var text string

	Count := 0

	for d := range str {
		Count++

		if str[d] == ' ' {
			break
		}
	}
	var buffer []string = make([]string, Count)
	for i := 0; i < Count+1; i++ {

		if str[i] == ' ' {

			text = str[:lol]
			buffer[i] = text
			text = ""
		}
	}

	return buffer

}
