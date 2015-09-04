package main

import "fmt"

func main() {

	//test := fmt.Sprintf("%s", "Hello there friend")
	//fmt.Println(test)

	age := fmt.Sprint("Hello, my name is Andy and I am ", 23, " years old.")
	fmt.Println(age)
	fmt.Printf("%T\n", age)

}
