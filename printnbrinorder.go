package student

import "github.com/01-edu/z01"

func PrintNbInOrder(n int64) {

	if n < 0 {

		return

	} else if n == 0 {

		z01.PrintRune('0')

	}

	var importantMassive []int64

	for n != 0 {

		C := n % 10

		n = n / 10

		importantMassive = append([]int64{C}, importantMassive...)
	}

	Size := 0

	for range importantMassive {

		Size++

	}

	for i := 0; i < Size; i++ {

		for x := 0; x < Size; x++ {

			if importantMassive[i] < importantMassive[x] {

				importantMassive[x], importantMassive[i] = importantMassive[i], importantMassive[x]

			}
		}
	}

	for x := range importantMassive {
		for y := '0'; y <= '9'; y++ {
			if rune(importantMassive[x]+48) == y {
				z01.PrintRune(y)
			}

		}

	}

}
