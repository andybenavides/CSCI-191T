package main

import "fmt"

func main()  {
	var name string
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hello",name, "we've been expecting you.")
}