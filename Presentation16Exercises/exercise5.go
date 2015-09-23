package main

import "fmt"

func dogYears (name string, age int) (int, bool) {
	dogAge := age * 7

	if (age > 25) {
		return dogAge, true
	}else{
		return dogAge, false
	}
}

func main() {
	name := "Andy"
	age := 23

	ageDogs, old := dogYears(name, age)

	if(old){
		ageString := fmt.Sprintf("%s is %v in dog years and is old",
			name, ageDogs)
		fmt.Println(ageString)
	}else{
		ageString := fmt.Sprintf("%s is %v in dog years and is not old",
			name, ageDogs)
		fmt.Println(ageString)
	}
}
