package student

import "github.com/01-edu/z01"

func PrintComb(i) {

	  if i[0] < i[1] && i[1] < i[2] {
		  z01.PrintRune(rune(i[0]))  
	  	  z01.PrintRune(rune(i[1]))
		  z01.PrintRune(rune(i[2]))
		  z01.PrintRune(',')
		  z01.PrintRune(' ')
	} 

}
