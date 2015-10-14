package main

import (
	"fmt"
)

type Sports struct{
	violent bool
	popular int
}

type Football struct{
	Sports
	BallShape string
}

type Basketball struct{
	Sports
	BallShape string
}

func main() {
	f1 := Football{Sports{true, 10},"Oval"}
	b1 := Basketball{Sports{false, 9}, "Sphere"}

	sports := []interface{}{f1, b1}
	fmt.Println(sports)
}

// The interface in this example is a slice holding three different types
