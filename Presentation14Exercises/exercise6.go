package main

import (
	"fmt"
	"strconv"
)

func main() {
	userInput := "23"
	userAge, _ := strconv.Atoi(userInput)
	fmt.Print(userAge)
	fmt.Printf("%T\n", userAge)


	userInput2 := 23
	userAge2 := strconv.Itoa(userInput2)
	fmt.Print(userAge2)
	fmt.Printf("%T\n", userAge2)
}
