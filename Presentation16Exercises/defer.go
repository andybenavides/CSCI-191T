package main

import "fmt"

func Double(num int){
	fmt.Println(num*2)
}

func Triple(num int){
	fmt.Println(num*3)
}

func main() {
	defer Double(4)
	Triple(4)
}
