package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64 = 0.0
	fmt.Println("Enter a number:")
	fmt.Scanln(&x)

	fmt.Println(math.Ceil(x))
}
