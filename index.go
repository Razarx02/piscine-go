package student

func Index(s string, toFind string) int {

	for x, word:= range s {

		if toFind[0] == word {

			return x

		}

	}

	return -1

}
