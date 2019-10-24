package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	A := os.Args[0]
	cat := 0

	for range A {
		cat++
	}

	for i := 0; i < cat; i++ {

		z01.PrintRune(rune(A[i]))
	}

}
