package main

import "fmt"

func fibonacci(x int) int {
	if x == 1 {
		return 1
		fmt.Println("x == 1, return 1")
	} else if x == 0 {
		return 0
		fmt.Println("x == 0, return 0")
	}
	fmt.Println(x,"...so...", x-1, x-2)
	return fibonacci(x-1) + fibonacci(x-2)
}

func main() {
	fmt.Println(fibonacci(10))
}
