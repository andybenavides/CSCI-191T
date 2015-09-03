package main

import "fmt"

func main() {

	Sum(1, 2, 3, 4)

}

func Sum(num1 int, num2 ...int){
	sum := 0
	sum = num1

	for _, value := range num2{
		sum += value
	}
	fmt.Println(sum)
}
