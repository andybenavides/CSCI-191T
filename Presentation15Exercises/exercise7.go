package main

import "fmt"

func main() {

	myMap := map[string]string{
		"Andy": "Student",
		"Cassie": "Student",
		"John": "Mechanic",
		"Katherine": "Teacher",
		"Joey": "Prisoner",
	}

	myMap["Sal"] = "Student"
	myMap["Andy"] = "Technician"

	if val, exists := myMap["Joey"]; exists{
		delete(myMap, "Joey")
		fmt.Println("Val:", val)
		fmt.Println("exists:", exists)
		fmt.Println("----Deleted----")
	}else{
		fmt.Println("That key does not exist")
		fmt.Println("Val: ", val)
		fmt.Println("exists: ", exists)
	}

	for key, val := range myMap{
		fmt.Println(key, "-", val)
	}

	fmt.Println(myMap)
	fmt.Println(len(myMap))
}
