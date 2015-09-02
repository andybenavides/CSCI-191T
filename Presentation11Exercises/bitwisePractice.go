package main

import "fmt"

func main()  {
	number := 100

	fmt.Println("The number value is: ", number)

	number = 100 >> 2
	fmt.Println("The number of a bitwise shift of 2 to the right is: ", number)
}