package main

import "fmt"

func main() {

	type Customer struct{
		name string
		numOfitems int
		paymentType string
	}

	customer1 := Customer{"Andy", 10, "Credit"}

	fmt.Println(customer1.name, customer1.numOfitems, customer1.paymentType)

	customer1.paymentType = "Cash"
	fmt.Println(customer1.name, customer1.numOfitems, customer1.paymentType)

	customer2 := new(Customer)
	customer2.name = "Jefferey"
	fmt.Println(customer2.name)
}

/*
	You cannot initialize a struct with make

	You can initialize a struct with new but it returns a pointer
 */