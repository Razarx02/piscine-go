package student

func ConcatParams(args []string) string {

	lol := 0

	for range args {

		lol++
	}

	var str string
	for i := 0; i < lol; i++ {

		str = str + args[i]
		if i != lol-1 {
			str = str + "\n"
		}

	}

	return str
}
