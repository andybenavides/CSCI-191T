package main

import "fmt"

type food struct{
	name string
	cost int
}

func (f *food) changeCost(newCost int) int{
	f.cost = newCost
	return f.cost
}

func main() {
	banana := food{"Banana", 2}
	fmt.Println(banana.cost)
	banana.changeCost(3)
	fmt.Println(banana.cost)
}

// Use a pointer to a food object because the nature of this particular
// struct is to be mutable.
// Using a value for the receiver would create a copy and 2 would print twice.