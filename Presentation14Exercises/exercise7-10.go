package main

import "fmt"

func main() {
	// ---------------- 7 ------------------//
	var myInt int = 23
	var myFloat64 float64 = float64(myInt)
	fmt.Println(myFloat64)

	// ---------------- 8 ------------------//
	var myIntFloat64 float64 = 100
	var myInt2 int = int(myIntFloat64)
	fmt.Println(myInt2)

	// ---------------- 9 ------------------//
	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'}))

	// ---------------- 10 -----------------//
	fmt.Println([]byte("My name is Andy"))
}


