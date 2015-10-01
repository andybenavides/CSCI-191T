package main

import "fmt"

type University struct{
	name string
	Enrollment int
}

type UniClass struct{
	University
	name string
	dayOfWeek string
}

func main() {
	class := UniClass{University{"Fresno State", 25000}, "Intro to math",
		"Monday"}

	fmt.Println(class)
}
