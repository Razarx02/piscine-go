package student

import "github.com/01-edu/z01"

func PrintStr(str string) {

	for index, word := range str {

		z01.PrintRune(rune(str[index]))

	}

	z01.PrintRune(10)

}
