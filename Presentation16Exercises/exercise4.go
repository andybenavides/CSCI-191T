package main

import "fmt"

func Greeting(name string, age int) string{
	greeting_string := fmt.Sprintf("%s is %v years old", name, age)
	return greeting_string
}

func main() {
	fmt.Println(Greeting("Andy", 24))
}
