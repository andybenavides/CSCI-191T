package main

import "fmt"

func main() {
	var myString string = "This is my string."
	var myRune rune = myString[0]
	fmt.Println(myRune)
	var mySlice string = myString[6:]
	fmt.Println(mySlice)
}
