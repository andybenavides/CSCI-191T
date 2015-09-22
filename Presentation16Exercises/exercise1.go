package main

import "fmt"

func half(int1 int) (int, bool){
	return int1/2, int1%2 == 0
}

func main() {

	x := 24

	fmt.Println(half(x))

}
