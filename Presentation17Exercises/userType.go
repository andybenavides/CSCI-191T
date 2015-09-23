package main

import "fmt"

type myType int

func main() {

	myType := 24

	fmt.Println(myType)
	fmt.Printf("%T\n", myType)
}

// underlying type for myType is int