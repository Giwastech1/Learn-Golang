package main

import "fmt"

func printStr(s string) {
	for _, str := range s {
		fmt.Print(string(str))
	}
	fmt.Print("\n")
}

func main() {
	printStr("Hello everyone!")
	printStr("Hello world!")
	printStr("Hello !")
}
