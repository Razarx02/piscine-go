package student

func Compare(a, b string) int {
	Cat := 0
	Dog := 0

	for range a {
		Cat++
	}
	for range b {
		Dog++
	}

	Z := Cat - Dog

	if b == a {

		return 0

	} else if b == a[Z:] {

		return -1

	} else {

		return 1
	}

}
