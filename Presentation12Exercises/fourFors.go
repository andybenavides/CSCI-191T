package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++  {
		fmt.Println(i)
	}

	fmt.Println("--------")

	value := 0
	for value <= 10{
		fmt.Println(value)
		value++
	}

	fmt.Println("--------")

	i := 110
	for{
		if i > 130 {break}
		fmt.Println(i)
		i++
	}
}
