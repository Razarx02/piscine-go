package main

import "github.com/01-edu/z01"

func main() {
	for i := 0; i < 10; i++ {
		z01.PrintRune(rune("%d", i))
	}
	z01.PrintRune()
}
