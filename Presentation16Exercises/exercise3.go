package main

import "fmt"

func longStatement() bool{
	return (true && false) || (false && true) || !(false && false)
}

func main() {
	fmt.Println(longStatement())
	// true
}
