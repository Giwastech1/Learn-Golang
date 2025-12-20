package main

import "fmt"

func strRev(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i <= length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(strRev("Hello world!"))
	fmt.Println(strRev("ghfghvvffffff"))
	fmt.Println(strRev("aé中"))
	fmt.Println(strRev("world"))

}