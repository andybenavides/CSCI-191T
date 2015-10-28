package main

import(
	"fmt"
	"reflect"
)

func main() {
	var age = "23"
	fmt.Println(reflect.TypeOf(age))

	fmt.Printf("%T\n", age)

}
