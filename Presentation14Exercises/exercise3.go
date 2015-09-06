package main

import "fmt"

func main() {
	rune := "This is my string"
	fmt.Println("The UTF-8 code for the last byte in that string is:", +
	rune[16])
}
