package student

import "github.com/01-edu/z01"

func PrintStr(str string) {

	a := len(str) - 1

	for i := 0; i <= a; i++ {
		z01.PrintRune(rune(str[i]))

	}

	z01.PrintRune(10)

}
