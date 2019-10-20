package student

func BasicAtoi(word string) string {

	toText := []byte(word)
	index := 0

	for range toText {

		index++
	}

	if toText[index-1] == '0' {

		return string('0')
	}

	for range []byte(word) {

		if toText[0] == '0' {

			toText = toText[1:]

		}

	}

	return string(toText)
}
