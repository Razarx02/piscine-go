package student

import "github.com/01-edu/z01"

func PrintNbIonroder(n int64) {

	if n < 0 {

		return

	} else if n == 0 {

		z01.PrintRune('0')

	}

	var importantMassive []int64
	importantMassive = IntoMassive(n, importantMassive)
	SortMassive(importantMassive)

	for x := range importantMassive {
		for y := '0'; y <= '9'; y++ {
			if rune(importantMassive[x]+48) == y {
				z01.PrintRune(y)
			}

		}

	}

}

func IntoMassive(N int64, Massive []int64) []int64 {
	for N != 0 {
		C := N % 10
		N = N / 10
		Massive = append([]int64{C}, Massive...)
	}
	return Massive
}

func SortMassive(Massive []int64) []int64 {
	Size := 0
	for range Massive {
		Size++
	}
	for i := 0; i < Size; i++ {
		for x := 0; x < Size; x++ {
			if Massive[i] < Massive[x] {
				Massive[x], Massive[i] = Massive[i], Massive[x]
			}
		}
	}
	return Massive
}

func main() {
   
	PrintNbIonroder(321)
}

