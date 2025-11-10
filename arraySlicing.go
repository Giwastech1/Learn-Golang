package main

import "fmt"

func sliceArray(){
	myArray := []int{20,30,45,7,43,4,99}
	fmt.Println(myArray)
	myNewArray := myArray[1:3]
	fmt.Println(myNewArray)
	fmt.Println(cap(myNewArray))
	myNewArray = myArray[1:5]
	fmt.Println(myNewArray)
}

func sliceWithMake(){
	myArray := make([] int, 8)
	fmt.Println(myArray)
}

func main(){
	sliceArray()
	sliceWithMake()
}