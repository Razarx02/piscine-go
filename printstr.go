package student

import "github.com/01-edu/z01"

func PrintStr(str string) {


	for i := 0; i <= len(str) -1 ; i++ {

		z01.PrintRune(rune(str[i]))

	}

	z01.PrintRune(10)

}
