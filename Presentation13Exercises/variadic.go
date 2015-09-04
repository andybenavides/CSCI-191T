package main

import "fmt"

func main() {

	greetings := []string{"Hello", "Hi"}

	printThese(greetings...)

}

func printThese(greetingToPrint ...string){
	fmt.Println(greetingToPrint)
}
