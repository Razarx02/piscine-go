package main

import (
	//"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	text := os.Args

	Cat := 0

	for range text {
		Cat++
	}

	for i := Cat - 1; i > 0; i-- {

		Buf := []rune(text[i])

		for x := range Buf {

			if Buf[x] == 32 {

				z01.PrintRune(10)
			} else {
				z01.PrintRune(Buf[x])
			}

		}

		z01.PrintRune(10)
	}

}
