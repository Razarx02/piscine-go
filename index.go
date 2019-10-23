package student

func Index(s string, toFind string) int {

	for x := range s {

		if toFind[0] == s[x] {

			return x

		}

	}

	return -1

}
