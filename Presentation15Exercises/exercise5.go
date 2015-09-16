package main

import "fmt"

func main() {

	sliceOfInts := []int{1,2,3,4}
	sliceOfInts2 := []int{5,6,7,8}

	// deleting 5
	sliceOfInts = append(sliceOfInts, sliceOfInts2[1:]...)

	fmt.Println(sliceOfInts)

}
