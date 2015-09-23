package main

import "fmt"

func main() {
	greeting := func(name string){
		fmt.Printf("Hello %s\n", name)
	}

	greeting("Andy")
	fmt.Printf("%T\n", greeting)
}
