package main

import "fmt"

func main() {
	double := func(x int) int{
		return x*2
	}

	fmt.Println(double(4))
}

// closed under func main
