package main

import "fmt"

type myType string

func Greeting (name myType){
	fmt.Println(name)
}

func main() {
	var name myType = "Andy"
	Greeting(name)
}
