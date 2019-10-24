package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	A := os.Args[0]

	d := len(A)
	for i := 0; i < d; i++ {

		z01.PrintRune(rune(A[i]))
	}

}
