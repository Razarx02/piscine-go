package main

import "github.com/01-edu/z01"

func main() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '8'; j++ {
			for Ewe := '0'; Ewe <= '9'; Ewe++ {
				for Ewe1 := j + 1; Ewe1 <= '9'; Ewe1++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(Ewe)
					z01.PrintRune(Ewe1)
					if i != '9' || j != '8' || Ewe != '9' || Ewe1 != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
