package main

import "fmt"

func main() {
	var age interface{}
	fmt.Printf("%T\n", age)
	age = "Andy"

	fmt.Printf("%T\n", age)
}
