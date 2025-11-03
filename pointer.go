package main

import "fmt"

func main(){
	var myAge = 16
	var pointer = &myAge


	fmt.Println(myAge)
	fmt.Println(pointer)
	fmt.Println(*pointer)

	*pointer = 29
	fmt.Println(*pointer)
}