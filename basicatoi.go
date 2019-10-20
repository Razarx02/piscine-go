package main

import "fmt"

func main() {

	a := "56465456455645"

	d := BasicAtoi(a)

	fmt.Println(d)


}

func BasicAtoi(word string) int {

	toText := []byte(word)
	Chek := true

	for range []byte(word) {

		if toText[0] == '0' {

			toText = toText[1:]
			Chek = false

		} else if toText[0] == ' ' {

			Chek = false

		} else {

			Chek = true
		}

	}

	if word == "" {

		Chek = false

	}

	if Chek == false {

		return '0'
	}

	Finaltext := int(toText)

	return int(Finaltext)
}
