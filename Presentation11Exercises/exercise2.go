package main

import "fmt"

const (
	_ = iota
	A = iota * .9
	B = iota * .8
	C = iota * .7
)

func main(){
	fmt.Println("You have been given a grade of: ", A)
}