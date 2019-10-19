package main

import "fmt" 




func main(){
	  
	a := 13
	b := 2


	DivMod(&a, &b)
	
	fmt.Println(a)
	fmt.Println(b)
}
func DivMod(a *int, b *int){

    x := *a
	
	*a = *a / *b
	*b =  x % *b

}