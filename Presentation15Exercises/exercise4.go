package main

import "fmt"

func main() {

	sliceOfInts := []int{1,2,3,4}
	sliceOfInts2 := []int{5,6,7,8}

	sliceOfInts = append(sliceOfInts, sliceOfInts2...)

	fmt.Println(sliceOfInts)

}
