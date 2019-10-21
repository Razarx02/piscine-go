package main

import "fmt"

func iterativeFactorial(s int) int {

	result := 0

	for i := 1; i > s+1; i++ {

		result = result + i

	}

	return result

}

func main() {

	b := iterativeFactorial(8)

	fmt.Println(b)

}
