package student

import "github.com/01-edu/z01"

func PrintComb() {
	
	for i := '0'; i <= '9'; i++ {
		
		for b := '0'; b <= '9'; b++ {
			for z := '0'; z <= '9'; z++{
		   		if i < b && b < z {
			   		z01.PrintRune(rune(i))
			   		z01.PrintRune(rune(b))
					z01.PrintRune(rune(z))   
					z01.PrintRune(',')
			   		z01.PrintRune(' ')
		   }
		   }
		}
	} 
	
	
}


