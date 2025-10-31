package main 

import "fmt"

func main(){
	mySlice := [...]int{}
	mySlice2 := []int{1,2,3}
	fmt.Println(mySlice)
	fmt.Println(mySlice2)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))


	mySlice3 := []string{"I", "want", "to", "write", "Golang"}
	fmt.Println(mySlice3)
	fmt.Println(len(mySlice3))
	fmt.Println(cap(mySlice3))
}