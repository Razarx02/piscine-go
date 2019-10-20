package student

func BasicAtoi(word string) string {

	toText := []byte(word)
	Chek := true

	for range []byte(word) {

		if toText[0] == '0' {

			toText = toText[1:]
			Chek = false

		} else {
			Chek = true
		}

	}

	if Chek == false {

		return string('0')
	}

	return string(toText)
}
