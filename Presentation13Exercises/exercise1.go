package main

import "fmt"

func main() {

	Sum(93451, 232, 223, 423)

	name := "Andy"
	fmt.Print(name, "\n")
	fmt.Printf("%v\n", name)
	fmt.Println(name)

}

func Sum(num2 ...int){
	sum := 0

	for _, value := range num2{
		sum += value
	}
	fmt.Println(sum)
}
