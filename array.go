package main

import "fmt"

func main() {
	var myArray = [3]int{1,2,3}
	array2:= [...]int{4,5,6,7,8,9,8}
	var array3 = [4]string{"book","pen","boy","girl"}
	fmt.Println(myArray)
	fmt.Println(array2)
	fmt.Println(array3[0],array3[2])
	// to change a value of an array
	array3[3] = "water" 
	fmt.Println(array3)

	// to find the length of an array, len () is used.
	array4 := [4]int {3,2,45,8}
	array5 := [...]int {1,2,4,6,7,54,66,77,566,45}
	fmt.Println(len(array4))
	fmt.Println(len(array5))
	fmt.Println(array2[2])
}