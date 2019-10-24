package student

func SplitWhiteSpaces(str string) []string {

	lol := 0

	for range str {
		lol++
	}

	var text string
	var buffer []string

	Count := 0

	for d := range str {
		Count++

		if str[d] == ' ' {
			break
		}
	}

	for i := 0; i < Count+1; i++ {

		if str[i] == ' ' {

			text = str[:lol]
			buffer = append(buffer, text)
			text = ""
		}
	}

	return buffer

}
