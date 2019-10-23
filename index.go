package student

func Index(s string, toFind string) int {
	D := 0

	for i, _ := range toFind {

		for x, _ := range s {

			if toFind[i] == s[x] {

				return D

			}

		}
		D--
	}

	return D

}
