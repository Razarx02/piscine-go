package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printStr(str string) {

	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	arguments := os.Args[1:]

	Cat := 0

	for range arguments {
		Cat++
	}

	if isEven(Cat) == true {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
