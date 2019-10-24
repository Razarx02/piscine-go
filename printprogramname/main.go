package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	text := []rune(os.Args[0])

	for i := range text {

		z01.PrintRune(text[i])

	}
	z01.PrintRune('\n')
}
