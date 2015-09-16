package main

import "fmt"

func main() {

	sliceofInts := make([]int, 5, 10)

	sliceofInts[0] = 1
	sliceofInts[1] = 2
	sliceofInts[2] = 3
	sliceofInts[3] = 4
	sliceofInts[4] = 5
	sliceofInts[5] = 6	// index out of range as length is set to 5 wherein
	// this is attempted to access the 6th element
	fmt.Println(sliceofInts)

}
