package main 

import "fmt"

func printBr(n int) {
	fmt.Print(n)
}


func main(){
	printBr(-123)
	printBr(0)
	printBr(123)
	fmt.Print("\n")
}