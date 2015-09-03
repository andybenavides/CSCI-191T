package main

import "fmt"

func main() {

	var num1, num2, rem = 0, 0, 0

	fmt.Println("Enter a number: ")
	fmt.Scan(&num1)
	fmt.Println("Enter another number: ")
	fmt.Scan(&num2)

	rem = num1 % num2

	fmt.Println("The remainder of ", num1, " and ", num2, " is: ", rem)
}
