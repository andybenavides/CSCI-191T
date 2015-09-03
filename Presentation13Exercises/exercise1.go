package main

import "fmt"

func main() {

	Sum(93451, 232, 223, 423)

}

func Sum(num2 ...int){
	sum := 0

	for _, value := range num2{
		sum += value
	}
	fmt.Println(sum)
}
