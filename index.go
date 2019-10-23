package student

func Index(s string, toFind string) int {

	D := 0

	for i := range toFind {

		for x := range s {

			if toFind[i] == s[x] {

				return D

			}

		}
		D--
	}

	return D

}
