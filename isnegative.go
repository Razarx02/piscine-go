package main

import "github.com/01-edu/z01"


func main(){
	var a = 5	
	Isnegative(a)   
   
}

func Isnegative(a int){
   
	if a >= 0{
		z01.PrintRune('T')
		z01.PrintRune(10)
	} else {
	 z01.PrintRune('F')
	 z01.PrintRune(10)
	}
         
   
 }
 

