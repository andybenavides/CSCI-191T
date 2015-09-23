package main

import "fmt"

func Average(x ...int) int{
	total := 0
	for _, num := range x{
		total = total + num
	}
	return total / len(x)
}

func main() {
	fmt.Println(Average(2,2,22))

}
