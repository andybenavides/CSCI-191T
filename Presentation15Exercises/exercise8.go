package main

import "fmt"

func main() {

	var myInt *int = new(int)
	fmt.Println(myInt)
	fmt.Println(*myInt)

	var myString *string = new(string)
	fmt.Println(myString)
	fmt.Println(*myString)

	var myBool *bool = new(bool)
	fmt.Println(myBool)
	fmt.Println(*myBool)

	var mySliceofInts []int = make([]int, 5)
	fmt.Println(mySliceofInts)

	var myMap map[int]string = make(map[int]string, 10)
	fmt.Println(myMap)

}

// new returns a pointer -- True
// make returns a pointer -- False

/*
	Zero values
	-----------
	Int: 0
	String: ""
	Bool: False
	Slice: [0]
	Map: map[]

	Difference between "make" and "new"
	-----------------------------------
	"make" will initialize a map, slice or a channel only and will allocate memory
	"new" returns a pointer to a zero valued allocated variable
 */