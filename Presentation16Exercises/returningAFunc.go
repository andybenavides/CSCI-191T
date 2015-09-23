package main

import "fmt"

func double(x int) func() string{
	return func() string{
		return "Nice number"
	}
}

func main() {
	function := double(4)
	fmt.Println(function())
}
