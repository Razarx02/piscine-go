package student

import "github.com/01-edu/z01"

func PrintComb2() {

	for j := '0'; j <= '9'; j++ {
		for i := '0'; i <= '9'; i++ {
			for Da := j; Da <= '9'; Da++ {

				for Has := '0'; Has <= '9'; Has++ {

					if ((j != i) || (Da != Has)) && ((j != Da) || (i != Has)) {

						if Has >= i {
							z01.PrintRune(j)
							z01.PrintRune(i)
							z01.PrintRune(' ')
							z01.PrintRune(Da)
							z01.PrintRune(Has)

							if j != '9' || i != '8' || Da != '9' || Has != '9' {
								z01.PrintRune(',')
								z01.PrintRune(' ')

							}
						}
					}

				}

			}

		}

	}

	z01.PrintRune('\n')

}
