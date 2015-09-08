package main

import "fmt"

func main() {

	var num1 = 23
	var string1 = string(num1)
	fmt.Printf(string1)

	var string2 = string('a')
	fmt.Println(string2)

	string3 := string([]byte{'h','e','l','l','o'})
	fmt.Println(string3)

	bytes := []byte("Hello")
	fmt.Println(bytes)

	integer1 := float64(12)
	fmt.Println(integer1)

	integer3 := 12.2
	fmt.Println(int(integer3))
}
