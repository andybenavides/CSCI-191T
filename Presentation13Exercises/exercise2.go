package main

import "fmt"

func main() {
	a := []int{23, 53, 543, 2}
	sum(a)
}

func sum(a []int){
	total := 0

	for _, value := range a{
		total += value
	}

	fmt.Println("The total of ", a, " is: ", total)
}