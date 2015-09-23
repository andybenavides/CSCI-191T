package main

import "fmt"

func double(number int, callback func(int)){
	callback(number)
}

func main() {
	double(14, func(x int){
		fmt.Println(x*2)
	})
}
