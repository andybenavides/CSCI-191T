package main

import "fmt"

func main() {
	SanFrancisco := 100
	var addressOf *int = &SanFrancisco

	fmt.Println("The address of the variable SanFrancisco is: ", addressOf)
	fmt.Println("The value of SanFrancisco is: ", SanFrancisco)

	*addressOf = 99
	fmt.Println("---------------------------------------------------")

	fmt.Println("The value of SanFrancisco is now: ", SanFrancisco)
	fmt.Println("The address of the variable SanFrancisco is still: ",
		addressOf)
}
