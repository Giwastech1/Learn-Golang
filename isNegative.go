package main

import "fmt"

func isNegative(nb int) {
	if nb < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

func main() {
	isNegative(1)
	isNegative(0)
	isNegative(-1)
}
