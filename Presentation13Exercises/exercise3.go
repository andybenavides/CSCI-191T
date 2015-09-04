package main

import "fmt"

func main() {

	age, birth := 0, 0

	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Println("You are", age, "years old.")

	birth = 2015 - age
	fmt.Println("You were probably born in the year", birth)
}
