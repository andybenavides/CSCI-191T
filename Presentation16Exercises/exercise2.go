package main

import "fmt"

func findGreatest (num ...int) int{
	var greatest int

	for _, x := range num{
		if x > greatest{
			greatest = x
		}
	}
	return greatest
}

func main() {

	greatest := findGreatest(23, 42, 42, 1, 199)
	fmt.Println(greatest)

}
