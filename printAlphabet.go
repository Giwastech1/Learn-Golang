package main

import "fmt"

func printAlphabet(){
	for str:= 'a'; str <= 'z'; str++ {
		fmt.Print(string(str))
	}
}