package main

import "fmt"

func main() {

	scores := "99 88 79"

	var student1, student2, student3 int

	fmt.Sscan(scores, &student1, &student2, &student3)

	fmt.Println("Student 1 score: ", student1)
	fmt.Println("Student 2 score: ", student2)
	fmt.Println("Student 3 score: ", student3)

}
